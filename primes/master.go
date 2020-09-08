package primes

import "math"

//master starts master process
func master(workers int, chs []chan int, In chan int) {
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

//broadcastMin broadcast the newest selected, global min
func broadcastMin(workers int, chs []chan int, next int) {
	for i := 0; i < workers; i++ {
		chs[i] <- next
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

func _min(a, b int) int  {
	if a <= b {
		return a
	}

	return b
}
