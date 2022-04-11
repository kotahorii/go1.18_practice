package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println(slices.BinarySearch([]int{1, 2, 3, 4, 5}, 3))
}

type Stack[T any] []T

func New[T any]() *Stack[T] {
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
