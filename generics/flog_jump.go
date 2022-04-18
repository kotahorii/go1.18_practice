package generics

import "math"

func FlogJump(h []int) int {
	dp := make([]int, len(h))
	for i := 1; i < len(h); i++ {
		if i == 1 {
			dp[i] = int(math.Abs(float64(h[i-1] - h[i])))
		} else if i >= 2 {
			v1 := dp[i-1] + int(math.Abs(float64(h[i-1]-h[i])))
			v2 := dp[i-2] + int(math.Abs(float64(h[i-2]-h[i])))
			dp[i] = int(math.Min(float64(v1), float64(v2)))
		}
	}
	return dp[len(h)-1]
}
