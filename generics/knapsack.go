package generics

import (
	"math"
	"sort"
)

func Knapsack(N, W int, w, v []int, dp [][]int) int {
	for i := 1; i < N; i++ {
		for j := 0; j < W; j++ {
			if j < w[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]),
					float64(dp[i-1][j-w[i-1]]+v[i-1])))
			}
		}
	}
	max := func(a []int) int {
		sort.Ints(a)
		return a[len(a)-1]
	}(dp[N-1])
	return max
}
