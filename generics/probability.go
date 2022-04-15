package generics

import (
	"golang.org/x/exp/slices"
)

func Combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func NumberOfSelectGoods(countMap map[int]int) int {
	return (countMap[100] * countMap[400]) +
		(countMap[200] * countMap[300])
}

func Count(a []int) map[int]int {
	res := map[int]int{}

	for _, x := range a {
		_, ok := res[x]
		if ok {
			res[x] += 1
		} else {
			res[x] = 1
		}
	}
	return res
}

func CombiOfCards(countMap map[int]int) int {
	return (countMap[1] * (countMap[1] - 1) / 2) +
		(countMap[2] * (countMap[2] - 1) / 2) +
		(countMap[3] * (countMap[3] - 1) / 2)
}

func CombiCardsV2(a []int) (res int) {
	for x := range a {
		_, ok := slices.BinarySearch(a, 100000-x)
		if ok {
			res += 1
		}
	}
	return
}
