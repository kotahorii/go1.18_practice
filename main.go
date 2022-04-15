package main

import (
	"fmt"
	"sort"

	"go1.18/generics"
	"golang.org/x/exp/constraints"
)

func main() {
	a := []int{1, 99999, 2, 99998}
	sort.Ints(a)
	fmt.Println(generics.CombiCardsV2(a))
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
