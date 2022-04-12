package generics

import "math"

func Gcf(x, y int) int {
	if y == 0 {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func IsPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n))+1); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
