package main

import (
	"fmt"

	"go1.18/generics"
)

func main() {
	xs := []int{0, 2, 1}
	ys := []int{0, 2, 2}
	fmt.Println(generics.Nearest(xs, ys))
}
