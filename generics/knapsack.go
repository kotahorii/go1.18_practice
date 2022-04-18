package generics

import (
	"fmt"
	"math"
	"sort"
)

func Make2DBoolArray(N, S int) [][]bool {
	dp := make([][]bool, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]bool, S+1)
	}
	dp[0][0] = true
	for j := 1; j < S+1; j++ {
		dp[0][j] = false
	}
	return dp
}

func Make2DIntArray(N, S int) [][]int {
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, S+1)
	}
	dp[0][0] = 0
	for j := 1; j < S+1; j++ {
		dp[0][j] = math.MinInt
	}
	return dp
}

// N, W := 4, 10
// w := []int{3, 6, 4, 2}
// v := []int{100, 210, 130, 57}
func Knapsack(N, W int, w, v []int, dp [][]int) int {
	for i := 1; i < N+1; i++ {
		for j := 0; j < W+1; j++ {
			if j < w[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]),
					float64(dp[i-1][j-w[i-1]]+v[i-1])))
			}
		}
	}
	fmt.Println(dp)
	max := func(a []int) int {
		sort.Ints(a)
		return a[len(a)-1]
	}(dp[N])
	return max
}

// a := []int{4, 1, 5}
// N, S := 3, 10
func SelectCards(N, S int, a []int, dp [][]bool) bool {
	for i := 1; i < N+1; i++ {
		for j := 0; j < S+1; j++ {
			if j < a[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				if dp[i-1][j] || dp[i-1][j-a[i-1]] {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			}
		}
	}
	return dp[N][S]
}
