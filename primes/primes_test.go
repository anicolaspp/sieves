package primes

import "testing"

func TestPrimes(t *testing.T) {

	primes := Primes(1000, 100)

	if len(primes) != 169 {
		t.Errorf("%v", len(primes))
	}
}

func benchmarkPrimes(n int, w int, b *testing.B) {

	for i := 0; i < b.N; i++ {
		Primes(n, w)
	}
}

func BenchmarkPrimes1(b *testing.B) {
	benchmarkPrimes(1000, 1, b)
}

func BenchmarkPrimes2(b *testing.B) {
	benchmarkPrimes(1000, 10, b)
}

func BenchmarkPrimes3(b *testing.B) {
	benchmarkPrimes(1000, 100, b)
}

func BenchmarkPrime4(b *testing.B) {
	benchmarkPrimes(100000, 1, b)
}

func BenchmarkPrimes5(b *testing.B) {
	benchmarkPrimes(100000, 10, b)
}

func BenchmarkPrimes6(b *testing.B) {
	benchmarkPrimes(100000, 100, b)
}
