package generics

import (
	"fmt"
	"math"
)

func Sqrt() float64 {
	r := 2.0
	a := 2.0

	for i := 1; i <= 3; i++ {
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

func BinarySqrt() float64 {
	l := 1.0
	r := 2.0
	for i := 1; i <= 18; i++ {
		m := (l + r) / 2
		if m*m < 2 {
			l = m
		} else {
			r = m
		}
		fmt.Printf("Step %d: m = %.12f\n", i, m)
	}
	return (l + r) / 2
}

func Q4_3_4() float64 {
	l := 1.0
	r := 2.0
	for i := 1; i <= 18; i++ {
		m := (l + r) / 2
		if math.Pow(m, 10) < 1000 {
			l = m
		} else {
			r = m
		}
		fmt.Printf("Step %d: m = %.12f\n", i, m)
	}
	return (l + r) / 2
}
