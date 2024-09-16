package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num <= 3 && num > 1 {
		return true
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func countDivisors(n int64) int {
	var count int
	var i int64 = 1
	for ; i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if i*i == n {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}

func countPrimeDivisors(l, r int64) int {
	res := 0
	var i int64 = l
	for ; i < r; i++ {
		count := countDivisors(i)
		if isPrime(count) {
			res++
		}
	}
	return res
}

func main() {
	var l, r int64
	fmt.Scan(&l, &r)

	fmt.Println(countPrimeDivisors(l, r))
}
