package generics

import (
	"math"
)

func Multiple(num, x, y int) (count int) {
	for i := 1; i < num+1; i++ {
		if i%x == 0 || i%y == 0 {
			count += 1
		}
	}
	return
}

func DoubleCard(n, s int) (count int) {
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if i+j > s {
				break
			}
			count += 1
		}
	}
	return
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func PrimeNum(n int) (result []int) {
	for i := 2; i < n+1; i++ {
		isPrime := true
		for j := 2; j < int(math.Sqrt((float64(i))))+1; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			result = append(result, i)
		}
	}
	return
}
