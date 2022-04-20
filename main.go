package main

import (
	"fmt"

	"go1.18/generics"
)

func main() {
	A, B, C := []int{3, 3}, []int{2, 1}, []int{6, 4}
	fmt.Println(generics.CalcDistance(A, B, C))
}
