// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

var casesGammaSign = []struct {
	Label string
	In    float64
	Out   int
}{
	{"", +inf, 1},
	{"", -inf, 1},
	{"", nan, 1},
	{"", 0.0, 1},
	{"", 1.0, 1},
	{"", -1.0, -1},
	{"", -2.0, 1},
}

func TestGammaSign(t *testing.T) { testutil.Test(t, tol, casesGammaSign, GammaSign) }

/*
func BenchmarkGammaSign(b *testing.B) {
	GlobalF = bench(b, cGammaSign, "",
		func(x []float64) float64 {
			return float64(GammaSign(x[0]))
		})
}
*/
