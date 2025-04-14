package special

import "math"

// removeCommonElements filters the slices a and b to remove any elements in both.
func removeCommonElements(a, b []float64) ([]float64, []float64, int, int) {
	na := len(a)
	nb := len(b)

	aa := make([]float64, na)
	copy(aa, a)
	bb := make([]float64, nb)
	copy(bb, b)

	for i := 0; i < na; i++ {
		for j := 0; j < nb; j++ {
			if a[i] == b[j] {
				aa = append(aa[:i], aa[i+1:]...)
				bb = append(bb[:j], bb[j+1:]...)
				i--
				j--
				na--
				nb--
				break
			}
		}
	}

	return aa, bb, na, nb
}

// powN1 returns (-1)**n
func powN1(n int) float64 {
	if n&1 == 0 {
		return 1
	} else {
		return -1
	}
}

// poly evaluates a polynomial cs[0] + cs[1].x + ... + cs[n].x^n using Horner's method
func poly(x float64, c ...float64) float64 {
	n := len(c)
	res := c[n-1]
	for i := n - 2; i >= 0; i-- {
		res = math.FMA(x, res, c[i])
	}
	return res

}
