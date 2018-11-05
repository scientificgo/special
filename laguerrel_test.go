// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutils"
)

var casesLaguerreL = []struct {
	Label    string
	In1      int
	In2, Out float64
}{
	{"", 2, nan, nan},
	{"", -2, 2, nan},
	{"", 0, 1.21, 1},
	{"", 1, -3.141, 4.141},
	{"", 2, 3.5, 0.125},
	{"", 22, 3.5, -0.46670027635922204483686005709132609588515316543499753499},
	{"", 41, -4.3255, 4.31280964460014077797090383609544382702962848439229792e+09},
}

func TestLaguerreL(t *testing.T) { testutils.Test(t, tol, casesLaguerreL, LaguerreL) }

/*
func BenchmarkLaguerreL(b *testing.B) {
	GlobalF = bench(b, cLaguerreL, "",
		func(x []float64) float64 {
			return LaguerreL(int(x[0]), x[1])
		})
}
*/
