package primes

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

	master(workers, chs, In)

	unfiltered := gatherWorkersData(workers, chs, result)

	primes := filterPrimes(unfiltered)

	return primes
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