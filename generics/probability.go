package generics

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
		if _, ok := res[x]; ok {
			res[x] += 1
		} else {
			res[x] = 1
		}
	}
	return res
}
