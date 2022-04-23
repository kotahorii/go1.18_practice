package generics

import (
	"fmt"
	"math"
)

func Eratosthenes(n int) {
	prime := make([]bool, n)
	for i := 2; i < n; i++ {
		prime[i] = true
	}

	for i := 2; i*i < n; i++ {
		if prime[i] {
			for x := 2 * i; x < n; x += i {
				prime[x] = false
			}
		}
	}

	for i := 0; i < n; i++ {
		if prime[i] {
			fmt.Printf("%d, ", i)
		}
	}
}

func Q4_4_2() float64 {
	n := 1000
	value := 0.0

	for i := 0; i < n; i++ {
		x := (2*float64(i) + 1) / (2 * float64(n))
		value += math.Pow(2, float64(x*x))
	}
	return value / float64(n)
}

func Q4_4_3(n int) (res int) {
	F := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			F[j] += 1
		}
	}
	fmt.Println(F)
	for i := 1; i <= n; i++ {
		res += i * F[i]
		fmt.Printf("i: %d, F[i]: %d\n", i, F[i])
	}
	return
}