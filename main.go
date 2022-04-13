package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(500000 * math.Log2(500000))
	fmt.Println(4000000 * math.Log2(4000000))
	fmt.Println(30000000 * math.Log2(30000000))
	fmt.Println(strings.Repeat("*", 25))

	fmt.Println(math.Pow(3000, 2))
	fmt.Println(math.Pow(10000, 2))
	fmt.Println(math.Pow(30000, 2))
	fmt.Println(strings.Repeat("*", 25))

	fmt.Println(math.Pow(2, 23))
	fmt.Println(math.Pow(2, 26))
	fmt.Println(math.Pow(2, 29))
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
