package main

import (
	"fmt"

	"go1.18/generics"
)

func main() {
	N := 5
	a := []int{2, 5, 3, 3, 1}
	dp1, dp2 := make([]int, N+1), make([]int, N+1)
	dp1[0], dp2[0] = 0, 0
	fmt.Println(generics.Study(N, a, dp1, dp2))
}
