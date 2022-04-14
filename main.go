package main

import (
	"fmt"

	"go1.18/generics"
	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(generics.PrimeFactorization(20211225))
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
