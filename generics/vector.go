package generics

import "math"

func CalcDistance(A, B, C []int) float64 {
	ax, ay := A[0], A[1]
	bx, by := B[0], B[1]
	cx, cy := C[0], C[1]

	BAx, BAy := ax-bx, ay-by
	BCx, BCy := cx-bx, cy-by
	CAx, CAy := ax-cx, ay-cy
	CBx, CBy := bx-cx, by-cy

	pattern := 2
	if BAx*BCx+BAy*BCy < 0 {
		pattern = 1
	} else if CAx*CBx+CAy*CBy < 0 {
		pattern = 3
	}

	if pattern == 1 {
		return math.Sqrt(math.Pow(float64(BAx), 2) +
			math.Pow(float64(BAy), 2))
	}
	if pattern == 3 {
		return math.Sqrt(math.Pow(float64(BCx), 2) +
			math.Pow(float64(BCy), 2))
	}
	S := math.Abs(
		float64(BAx)*float64(CAy) -
			float64(BAy)*float64(CAx),
	)
	BCLength := math.Sqrt(
		math.Pow(float64(BCx), 2) +
			math.Pow(float64(BCy), 2),
	)

	return S / BCLength
}
