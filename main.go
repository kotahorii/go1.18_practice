package main

import (
	"fmt"

	"go1.18/generics"
)

func main() {
	numbers := []int{4, 2, 5, 3, 6, 1}
	fmt.Println(generics.FlogJump(numbers))
}
