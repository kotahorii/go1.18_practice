package main

import (
	"fmt"

	"go1.18/generics"
)

func main() {
	i := []generics.MyInt{1, 2, 3, 4, 5}
	s := []generics.MyString{"a", "b", "c", "d", "e"}
	fmt.Println(generics.F(i))
	fmt.Println(generics.F(s))
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
