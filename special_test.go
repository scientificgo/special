package special_test

import (
	"math"
)

const tol = 2e-10

var (
	nan = math.NaN()
	inf = math.Inf(1)
)

func equalFloat64(x float64, y float64) bool {
	if math.IsNaN(y) {
		return math.IsNaN(x)
	}

	if math.IsInf(y, 0) {
		return x == y
	}

	if y == 0 {
		return math.Abs(x) < tol
	}

	return math.Abs((x-y)/y) < tol
}
