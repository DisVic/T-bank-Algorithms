package main

import (
	"fmt"
	"math"
	"sync"
)

func sieveOfEratosthenes(max int) []int {
	sieve := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		sieve[i] = true
	}

	for p := 2; p*p <= max; p++ {
		if sieve[p] {
			for multiple := p * p; multiple <= max; multiple += p {
				sieve[multiple] = false
			}
		}
	}

	primes := []int{}
	for i := 2; i <= max; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func isPrime(n int, primes []int) bool {
	if n < 2 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for _, prime := range primes {
		if prime > limit {
			break
		}
		if n%prime == 0 {
			return false
		}
	}
	return true
}

func countDivisorsInRange(l, r int64, primes []int) int {
	rangeSize := r - l + 1
	divisorsCount := make([]int, rangeSize)

	var i int64
	for i < rangeSize {
		divisorsCount[i] = 1
		i++
	}

	var wg sync.WaitGroup
	numWorkers := 12

	chunkSize := rangeSize / int64(numWorkers)

	worker := func(start, end int64) {
		defer wg.Done()
		for _, prime := range primes {
			startIdx := max(int64(prime*prime), (l+int64(prime)-1)/int64(prime)*int64(prime)) - l
			for j := startIdx; j <= end-l; j += int64(prime) {
				temp := l + j
				count := 0
				for temp%int64(prime) == 0 {
					temp /= int64(prime)
					count++
				}
				divisorsCount[j] *= (count + 1)
			}
		}
		for i := start; i <= end; i++ {
			if divisorsCount[i-l] == 1 {
				divisorsCount[i-l] = 2
			}
		}
	}

	for i := int64(0); i < int64(numWorkers); i++ {
		wg.Add(1)
		start := l + i*chunkSize
		end := min(l+(i+1)*chunkSize-1, r)
		go worker(start, end)
	}
	wg.Wait()

	count := 0
	for _, divCount := range divisorsCount {
		if divCount > 1 && isPrime(divCount, primes) {
			count++
		}
	}

	return count
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	l := int64(1)
	r := int64(1000000000)

	sqrtR := int(math.Sqrt(float64(r)))
	primes := sieveOfEratosthenes(sqrtR)

	result := countDivisorsInRange(l, r, primes)
	fmt.Println(result)
}
