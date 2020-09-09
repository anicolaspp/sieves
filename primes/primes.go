package primes

import (
	"github.com/anicolaspp/sieves/primes/worker"
	"github.com/anicolaspp/sieves/primes/master"
)

//Primes calculates the list of prime numbers less or equal to number using some workers
// It uses sieves of Eratosthenes parallel algorithm
func Primes(number int, workers int) []int {

	chs, In, result := startWorkers(number, workers)

	master.Master(workers, chs, In)

	primes := worker.GatherData(chs, result)

	return primes
}

func startWorkers(number int, workers int) ([]chan int, chan int, chan []int) {
	chs := make([]chan int, workers)
	In := make(chan int)
	result := make(chan []int)

	for i := 0; i < workers; i++ {
		chs[i] = make(chan int)

		data := worker.NewData(i, chs[i], In, result)
		go worker.Worker(i, number/workers, data)
	}

	return chs, In, result
}

