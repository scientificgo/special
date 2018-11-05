// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutils"
)

var casesLambertW = []struct {
	Label    string
	In1      int
	In2, Out float64
}{
	{"", 2, 2.2, nan},
	{"", 0, 0, 0},
	{"", 0, -3, nan},
	{"", 0, +inf, +inf},
	{"", -1, 0, -inf},
	{"", -1, -3, nan},
	{"", -1, 3, nan},
	{"", 0, -0.33, -0.6032666497551331},
	{"", 0, 0.1, 0.09127652716086226},
	{"", 0, 4.5, 1.2672378143074348},
	{"", 0, 9.9, 1.7391425517333516},
	{"", 0, 10.89, 1.8000374607381258},
	{"", 0, 100.12, 3.386555992882349},
	{"", -1, -0.36787944117144233, -1},
	{"", -1, -0.33, -1.541268224332639},
	{"", -1, -0.1, -3.577152063957297},
}

func TestLambertW(t *testing.T) {
	testutils.Test(t, tol, casesLambertW, LambertW)
}

/*
func BenchmarkLambertW(b *testing.B) {
	GlobalF = bench(b, cLambertW, "",
		func(x []float64) float64 {
			return LambertW(int(x[0]), x[1])
		})
}
*/
