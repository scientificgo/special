package special

import "math"

// GammaSign returns the sign of the Gamma function at x.
// For non-positive integer x, it is sign of the residue, i.e. (-1)**|x|.
func GammaSign(x float64) int {
	if x >= 0 || math.IsInf(x, -1) || math.IsNaN(x) {
		return 1
	}
	n := math.Abs(math.Floor(x))
	return powN1(int(n))
}
