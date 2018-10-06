// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

// removeCommonElements filters the slices a and b to remove any elements in both.
func removeCommonElements(a, b []float64) ([]float64, []float64, int, int) {
	na := len(a)
	nb := len(b)
	for i := 0; i < na; i++ {
		for j := 0; j < nb; j++ {
			if a[i] == b[j] {
				a = append(a[:i], a[i+1:]...)
				b = append(b[:j], b[j+1:]...)
				i--
				j--
				na--
				nb--
				break
			}
		}
	}
	return a, b, na, nb
}
