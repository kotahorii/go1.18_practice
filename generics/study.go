package generics

import "math"

func Study(N int, a, dp1, dp2 []int) int {
	for i := 1; i <= N; i++ {
		dp1[i] = dp2[i-1] + a[i-1]
		dp2[i] = int(math.Max(float64(dp1[i-1]), float64(dp2[i-1])))
	}
	return int(math.Max(float64(dp1[N]), float64(dp2[N])))
}
