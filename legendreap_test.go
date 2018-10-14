// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	. "scientificgo.org/special"
	"scientificgo.org/testutils"
	"testing"
)

var casesLegendreAP = []struct {
	Label    string
	In1, In2 int
	In3, Out float64
}{
	{"", 2, 10, 3.5, 0},
	{"", -10, 10, 3.5, 0},
	{"", 153, 11, 1.1, nan},
	{"", 3, 0, -0.9, -0.4725},
	{"", 10, -3, 0.98, 0.000127413538489068664787307858757370814480523438403071987},
	{"", 22, 10, 3.5, -6.309576310867162196764903515496706859266851097345352e+29},
	{"", 153, 11, 1e-5, 7.1211274720051580808679401818532702913164641927292374e+22},
	{"", -154, 11, 1e-5, 7.1211274720051580808679401818532702913164641927292374e+22},
}

func TestLegendreAP(t *testing.T) { testutils.Test(t, tol, LegendreAP, casesLegendreAP) }

/*
func BenchmarkLegendreAP(b *testing.B) {
	GlobalF = bench(b, cLegendreAP, "",
		func(x []float64) float64 {
			return LegendreAP(int(x[0]), int(x[1]), x[2])
		})
}
*/
