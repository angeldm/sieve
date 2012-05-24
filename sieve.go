package sieve

import (
	"fmt"
	"sort"
)

type Sieve struct {
	Primes []int
	table  []byte
	i      int
}

func NewSieve(upperBound int) *Sieve {
	s := &Sieve{}
	s.Primes = make([]int, 0, upperBound/3)
	s.table = make([]byte, upperBound+1)
	for i := 2; i < len(s.table); i++ {
		if s.table[i] == 0 {
			// strike multiples from table 
			for j := i + i; j < len(s.table); j += i {
				s.table[j] = 1
			}
		}
	}
	for i := 2; i < len(s.table); i++ {
		if s.table[i] == 0 {
			s.Primes = append(s.Primes, i)
		}
	}

	return s
}

func (s Sieve) Chan() chan int {
	out := make(chan int, 4)
	go func() {
		for i := 0; i < len(s.Primes); i++ {
			out <- s.Primes[i]
		}
		out <- 0
	}()
	return out
}

func (s Sieve) ChanDescending(n int) chan int {
	out := make(chan int, 4)
	i := sort.SearchInts(s.Primes, n)
	if s.Primes[i] != n {
		i--
	}
	go func() {
		for ; i >= 0; i-- {
			out <- s.Primes[i]
		}
		out <- 0
	}()
	return out
}

func (s *Sieve) print() {
	for i, mask := range s.table {
		if i > 1 && mask == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")
}

func (s *Sieve) count() (tally int) {
	for i, mask := range s.table {
		if i > 1 && mask == 0 {
			tally++
		}
	}
	return
}

func (s *Sieve) test(n int) (isPrime bool) {
	if n > 1 && n < len(s.table) && s.table[n] == 0 {
		isPrime = true
	}
	return
}
