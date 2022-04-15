package generics

import "math"

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
