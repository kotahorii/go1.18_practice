package main

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(math.Sqrt(float64(841)))
	fmt.Println(math.Pow(29, 2))

	fmt.Println(math.Round(math.Pow(1024, 1/5.0)))
	fmt.Println(math.Pow(4, 5))
}

type Stack[T constraints.Ordered] []T

func New[T constraints.Ordered]() *Stack[T] {
	v := make(Stack[T], 0)
	return &v
}

func (s *Stack[T]) Push(x T) {
	(*s) = append((*s), x)
}

func (s *Stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}
