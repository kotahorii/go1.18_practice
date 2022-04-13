package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(13 & 14)
	fmt.Println(13 | 14)
	fmt.Println(13 ^ 14)

	fmt.Println(8 | 4 | 2 | 1)
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
