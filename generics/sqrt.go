package generics

import (
	"fmt"
	"math"
)

func Sqrt() float64 {
	r := 2.0
	a := 2.0

	for i := 1; i <= 5; i++ {
		x := a
		y := a * a

		tangent_a := 2.0 * x
		tangent_b := y - tangent_a*x

		next_a := (r - tangent_b) / tangent_a
		fmt.Printf("Step %d: a = %.12f -> %.12f\n", i, a, next_a)
		a = next_a
	}
	return a
}

func CubeRoot2() float64 {
	r := 2.0
	a := 3.0

	for i := 1; i < 10; i++ {
		x := a
		y := math.Pow(a, 3)

		tangent_a := 3 * x * x
		tangent_b := y - tangent_a*x

		next_a := (r - tangent_b) / tangent_a
		fmt.Printf("Step %d: a = %.12f -> %.12f\n", i, a, next_a)
		a = next_a
	}
	return a
}

func BinarySart()
