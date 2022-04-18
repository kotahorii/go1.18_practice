package main

import (
	"fmt"
	"math"

	"go1.18/generics"
)

func main() {
	N, W := 4, 10
	w := []int{3, 6, 4, 2}
	v := []int{100, 210, 130, 57}
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, W)
	}
	for i := 1; i < W; i++ {
		dp[0][i] = math.MinInt
	}

	fmt.Println(generics.Knapsack(N, W, w, v, dp))
}
