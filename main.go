package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := New[string]()
	s.Push("hello")
	s.Push("hello")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

// func f[T Stringer](xs []T) []string {
// 	var result []string
// 	for _, x := range xs {
// 		result = append(result, x.String())
// 	}
// 	return result
// }

type Stringer interface {
	String() string
}

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
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
