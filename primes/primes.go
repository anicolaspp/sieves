package primes

import (
	"math"
)

//Primes calculates the list of prime numbers less or equal to number using some workers
// It uses sieves of Eratosthenes parallel algorithm
func Primes(number int, workers int) []int {

	chs := make([]chan int, workers)
	In := make(chan int)
	result := make(chan []int)

	for i := 0; i < workers; i++ {
		chs[i] = make(chan int)

		data := &workerData{wid: i, data: []int{}, In: chs[i], out: In, result: result}
		go worker(i, number/workers, data)
	}

	runMaster(workers, chs, In)

	unfiltered := gatherWorkersData(workers, chs, result)

	primes := filterPrimes(unfiltered)

	return primes
}

//gatherWorkersData sends a signal to each worker so they send their local primes to the corresponding channel
func gatherWorkersData(workers int, chs []chan int, result chan []int) []int {
	unfiltered := make([]int, 0)

	for i := 0; i < workers; i++ {
		// send signal to gather worker data
		chs[i] <- -1

		// gather worker primes
		unfiltered = append(unfiltered, <-result...)
	}

	return unfiltered
}

//filterPrimes removes unnecessary values from the result (-1)
func filterPrimes(unfiltered []int) []int {
	primes := make([]int, 0)

	for _, v := range unfiltered {
		if v > -1 {
			primes = append(primes, v)
		}
	}
	return primes
}

func runMaster(workers int, chs []chan int, In chan int) {
	next := 2

	for {
		broadcastMin(workers, chs, next)

		min, hasMin := agreeOnNextMin(workers, In)

		if !hasMin {
			break
		}

		next = min
	}
}

//agreeOnNextMin collects the local min of each partition/worker and select to global min
func agreeOnNextMin(workers int, In chan int) (int, bool) {
	min := math.MaxInt32

	for i := 0; i < workers; i++ {
		localMin := <-In

		min = _min(min, localMin)
	}

	if min == math.MaxInt32 {
		return -1, false
	}

	return min, true
}

//broadcastMin broadcast the newest selected, global min
func broadcastMin(workers int, chs []chan int, next int) {
	for i := 0; i < workers; i++ {
		chs[i] <- next
	}
}

func _min(a, b int) int  {
	if a <= b {
		return a
	}

	return b
}
