// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

const tolHyp2F1 = 5 * Macheps

var casesHyp2F1 = []struct {
	Label                   string
	In1, In2, In3, In4, Out float64
}{
	{"", 1, 2, 5, 0.5, 1.2710646668773748520271418},
	{"", 1, 2, 5, -0.5, 0.83907329727298698549833412},
	{"", 20, 11, 3, 0.25, 548991.77233826151912145437},
	{"", 20, 11, 3, 0.5, 5.6302076632177777777777778E+11},
	{"", 20, 11, 3, 0.75, 8.8358430723507067289600000E+20},
	{"", 200, 101, 2, 0.75, 1.0791540760396911473815479E+251},
	{"", 200, 101, 400, 0.75, 6.4483651342221470083789865E+22},

	{"", 2000, 0.25, 0.125, 0.25, 3.53207513510593459216394617E+250},

	{"", 1, -2, 5, -0.5, 1.2166666666666666666666667},

	// special cases
	{"sc", 1, 1e100, 1e100, 0.1, 1.1111111111111111111111111}, // Hyp2F1(1, a, a, x) = 1/(1-x) for |x| < 1
	{"sc", 1, 2, 2, 0.1, 1.1111111111111111111111111},         // same as above
	{"sc", 2e6, 1, 2e6, 0.2, 1.25},                            // same as above

	{"sc", 1, 2, -3, 0, 1},   // Hyp2F1(a, b, c, 0) = 1
	{"sc", 1, 2, 3, 4, NaN},  // Hyp2F1(a, b, c, x > 1) = NaN
	{"sc", 1, 2, 3, -4, NaN}, // Hyp2F1(a, b, c, x < -1) = NaN
}

func TestHyp2F1(t *testing.T) { testutil.Test(t, tolHyp2F1, casesHyp2F1, Hyp2F1) }

func BenchmarkHyp2F1(b *testing.B) {
	for _, c := range casesHyp2F1 {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Hyp2F1(c.In1, c.In2, c.In3, c.In4)
			}
		})
	}
}
