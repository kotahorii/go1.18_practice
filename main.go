package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	cnt := 0
	for i := 1; i <= 5; i++ {
		for j := i + 1; j <= 5; j++ {
			cnt++
		}
	}
	fmt.Println(cnt)
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
