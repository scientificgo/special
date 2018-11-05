// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutils"
)

var casesHermiteH = []struct {
	Label    string
	In1      int
	In2, Out float64
}{
	{"", 2, nan, nan},
	{"", -2, 2, nan},
	{"", 0, 1.1111, 1},
	{"", 1, 1.1111, 2.2222},
	{"", 4, 5.5, 13201},
	{"", 43, 53.5, 1.56492249523929575819638209988328890215202268133174998288e+87},
}

func TestHermiteH(t *testing.T) { testutils.Test(t, tol, casesHermiteH, HermiteH) }

/*
func BenchmarkHermiteH(b *testing.B) {
	GlobalF = bench(b, cHermiteH, "",
		func(x []float64) float64 {
			return HermiteH(int(x[0]), x[1])
		})
}
*/
