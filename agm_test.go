// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"math"
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

var tolAGM = 1 * Macheps

var casesAGM = []struct {
	Label         string
	In1, In2, Out float64
}{
	{"", 1, 5, 2.6040081905309402886964274},
	{"", 20, 25, 22.430285802876025701278022},
	{"", 0.5, 1000, 174.78155546038950365420555},
	{"", 1e10, 1e15, 121774521865389.04490416355},
	{"", 1e-10, 1e10, 331126196.70463757381569721},
	{"", 1e200, 1e305, 6.4599892935511101190071753e+302},
	{"", 1e306, 1e305, 4.2504070949322748617281643e+305},

	// special cases
	{"sc", 0, 0, 0},
	{"sc", 1, 1, 1},
	{"sc", 0, -1, 0},
	{"sc", +Inf, +Inf, +Inf},
	{"sc", math.Pi, math.Pi, math.Pi},
	{"sc", +Inf, 0, NaN},
	{"sc", 1, -1, NaN},
	{"sc", NaN, 1, NaN},
}

func TestAGM(t *testing.T) { testutil.Test(t, tolAGM, casesAGM, AGM) }

func BenchmarkAGM(b *testing.B) {
	for _, c := range casesAGM {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = AGM(c.In1, c.In2)
			}
		})
	}
}
