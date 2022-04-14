package generics

import (
	"math"
)

func Divisors(n int) (result []int) {
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i != 0 {
			continue
		}
		result = append(result, i)
		if i != n/i {
			result = append(result, n/i)
		}
	}
	return
}

func PrimeFactorization(n int) (result []int) {
	v := n
	for i := 2; i <= int(math.Sqrt(float64(v))); {
		if n%i == 0 {
			result = append(result, i)
			n /= i
			continue
		}
		i++
	}
	return
}
