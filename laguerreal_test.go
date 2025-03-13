// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "github.com/scientificgo/special"
)

var casesLaguerreAL = []struct {
	Label         string
	In1           int
	In2, In3, Out float64
}{
	{"", 2, 1, nan, nan},
	{"", -2, 1, 2, nan},
	{"", 0, 22, 9.98, 1},
	{"", 1, 1, -3.141, 5.141},
	{"", 41, -47, 4.3255, nan},
	{"", 2, -1, 3.5, 2.625},
	{"", 22, 10, 3.5, 25252.39978746937700539858666621210279857995048910207458046},
	{"", 41, 17, -4.3255, 4.2086917252857569297940591579340234265295868262479278e+17},
	{"", 41, -17, 4.3255, 4.9579126998422110619614409645401386382056238641622942e-08},
}

func TestLaguerreAL(t *testing.T) {
	for i, c := range casesLaguerreAL {
		t.Run(c.Label, func(tt *testing.T) {
			res := LaguerreAL(c.In1, c.In2, c.In3)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("[%v]: Got %v, want %v", i, res, c.Out)
			}
		})
	}
}

/*
func BenchmarkLaguerreAL(b *testing.B) {
	GlobalF = bench(b, cLaguerreAL, "",
		func(x []float64) float64 {
			return LaguerreAL(int(x[0]), x[1], x[2])
		})
}
*/
