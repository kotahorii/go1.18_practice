package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(f([]MyInt{1, 2, 3, 4}))
}

func f[T Stringer](xs []T) []string {
	var result []string
	for _, x := range xs {
		result = append(result, x.String())
	}
	return result
}

type Stringer interface {
	String() string
}

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}
