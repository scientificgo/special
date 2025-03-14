package special

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

// s = (-1)**n
func minus1pow(n int) float64 {
	if n%2 == 0 {
		return 1
	} else {
		return -1
	}
}
