package generics

import (
	"math"
	"math/rand"
)

func CountPointInTheCircle(n int) (count int) {
	for i := 0; i <= n; i++ {
		px := 6 * rand.Intn(1000000) / 1000000
		py := 9 * rand.Intn(1000000) / 1000000

		dist_33 := math.Sqrt(float64(math.Pow((float64(px-3)), 2) +
			math.Pow(float64(py-3), 2)))
		dist_37 := math.Sqrt(float64(math.Pow((float64(px-3)), 2) +
			math.Pow(float64(py-7), 2)))

		if dist_33 <= 3.0 || dist_37 <= 2.0 {
			count++
		}
	}
	return
}
