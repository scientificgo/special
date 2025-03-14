package special

import "math"

// GammaSign returns the sign of the Gamma function at x. For non-positive integer x,
// the sign is given by the sign of the residue, i.e. (-1)**|x|.
func GammaSign(x float64) int {
	if x >= 0 {
		return 1
	}
	if x != math.Trunc(x) {
		x = math.Trunc(x - 1)
	}
	return 1 - 2*(int(x)&1)
}
