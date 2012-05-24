package sieve

import (
	"testing"
)

func TestCount(t *testing.T) {
	sieve := NewSieve(1000000000)
	count := sieve.count()
	if count != 50847534 {
		t.Errorf("Incorrect count %d of primes up to 1000000000, expected 50847534\n", count)
	}
}

func BenchmarkSieve(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	sieve := NewSieve(1000000000)
	sieve.count()
}
