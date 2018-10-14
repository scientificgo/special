// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	. "scientificgo.org/special"
	"scientificgo.org/testutils"
	"testing"
)

var casesSphericalHarmonicY = []struct {
	Label                string
	In1, In2             int
	In3, In4, Out1, Out2 float64
}{
	{"", 2, 1, 1, nan, nan, nan},
	{"", -22, -20, 10, 3.5, nan, nan},
	{"", 0, 0, -7.21, 7.11, 0.282094791773878143474039725780386292922025314664499428422, 0},
	{"", 22, -20, 10, 3.5, 0.00004474740985105335681955012220070375819491035216485506, -0.00005467954127843340910384743380389943632078690198859652},
	{"", 31, 31, -10, 3.5, 5.20105344181401315495645487292339772715804207504464e-10, -4.50058978347382319639863845184135153633008312108735e-09},
	{"", 7, -7, -7, 7, -0.0079410539609713541755773378737370761151396061682892738, -0.025196238025923628008773731309078469019874758134402119},
}

func TestSphericalHarmonicY(t *testing.T) {
	testutils.Test(t, tol, SphericalHarmonicY, casesSphericalHarmonicY)
}

/*
func BenchmarkSphericalHarmonicY(b *testing.B) {
	GlobalF = bench(b, cSphericalHarmonicY, "",
		func(x []float64) float64 {
			r, _ := SphericalHarmonicY(int(x[0]), int(x[1]), x[2], x[3])
			return r
		})
}
*/
