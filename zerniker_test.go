// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	. "scientificgo.org/special"
	"scientificgo.org/testutils"
	"testing"
)

var casesZernikeR = []struct {
	Label    string
	In1, In2 int
	In3, Out float64
}{
	{"", -1, 20, 1.1, nan},
	{"", 987, 988, 98765432.1234567, 0},
	{"", 7, 5, 3.3, 27484.7865039},
	{"", 43, 41, 53.5, 8.9646400010525796235939120335742702428482112823125695e+75},
}

func TestZernikeR(t *testing.T) { testutils.Test(t, tol, ZernikeR, casesZernikeR) }

/*
func BenchmarkZernikeR(b *testing.B) {
	GlobalF = bench(b, cZernikeR, "",
		func(x []float64) float64 {
			return ZernikeR(int(x[0]), int(x[1]), x[2])
		})
}
*/
