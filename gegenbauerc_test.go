// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	. "scientificgo.org/special"
	"scientificgo.org/testutils"
	"testing"
)

var casesGegenbauerC = []struct {
	Label         string
	In1           int
	In2, In3, Out float64
}{
	{"", 2, 1, nan, nan},
	{"", -2, 1, 2, nan},
	{"", 0, 3.3, 1, 1},
	{"", 1, 1, -3.141, -6.282},
	{"", 1, 2, 3, 12},
	{"", 2, 10, 3.5, 2685},
	{"", 22, 10, 3.5, 5.7734358481154896492250325e+25},
	{"", 40, 15, 4.3255, 6.6613993107922727134428780696365820449481333572147195e+49},
	{"", 40, -15, 4.3255, 0},
	{"", 40, -14.9, 4.3255, -1.280764317783101302989751939635266543812346402681405e+24},
}

func TestGegenbauerC(t *testing.T) { testutils.Test(t, tol, GegenbauerC, casesGegenbauerC) }

/*
func BenchmarkGegenbauerC(b *testing.B) {
	GlobalF = bench(b, cGegenbauerC, "",
		func(x []float64) float64 {
			return GegenbauerC(int(x[0]), x[1], x[2])
		})
}
*/
