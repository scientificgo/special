// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package poly_test

import (
	. "scientificgo.org/special/poly"
	"scientificgo.org/testutils"
	"testing"
)

var casesLegendreP = []struct {
	Label    string
	In1      int
	In2, Out float64
}{
	{"", 2, nan, nan},
	{"", 0, 1.21, 1},
	{"", 1, -3.141, -3.141},
	{"", 153, 1e-5, 0.000098854224351825737394494825999239936670301828815467202},
	{"", -2, -3.141, -3.141},
}

func TestLegendreP(t *testing.T) { testutils.Test(t, tol, LegendreP, casesLegendreP) }

/*
func BenchmarkLegendreP(b *testing.B) {
	GlobalF = bench(b, cLegendreP, "",
		func(x []float64) float64 {
			return LegendreP(int(x[0]), x[1])
		})
}
*/
