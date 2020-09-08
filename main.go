package main

import (
	"fmt"
	"github.com/anicolaspp/sieves/primes"
	"os"
	"strconv"
	"time"
)

func main() {

	number, _ := strconv.Atoi(os.Args[1])
	workers, _ := strconv.Atoi(os.Args[2])

	start := time.Now()
	somePrimes := primes.Primes(number, workers)

	end := time.Now().Nanosecond() - start.Nanosecond()

	fmt.Println(nanoToSecond(end))
	fmt.Println(somePrimes)
}

func nanoToSecond(nano int) float64  {
	return float64(nano) / 1000000000
}

