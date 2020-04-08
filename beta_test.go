// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

const tolBeta = 6 * Macheps

var casesBeta = []struct {
	Label         string
	In1, In2, Out float64
}{
	// direct calculation
	{"", 1, 10, 0.1},
	{"", 19.125, 11.0125, 2.4472541189858992560452364E-9},
	{"", 39, 10, 1.5288846296038539937830683E-11},
	{"", 8.25, -40.75, 1.3524462378721093720166038E-9},

	// argument reduction
	{"", 80, 100, 7.480703996850429143482604E-55},
	{"", 171, 171, 3.0280694105736877516074272E-104},
	{"", 240, 230, 8.4432089336112537679900408E-143},
	{"", 100., -98.5, -35.315948806233095764437743},
	{"", 100, -101.5, 1.3048090539541725894540623E-5},
	{"", 240, -30.5, -1.0391711818066296050416532E+39},

	// special cases
	{"sc", 1, -2, -0.5},          // Beta(1, -2) = -1/2
	{"sc", 10, -10, 0.1},         // Beta(10, -10) = 1/10
	{"sc", 101, -101, -1. / 101}, // Beta(x, -x) = (-1)**x / x
	{"sc", -3.5, 2.5, 0},         // Beta(x, y) = 0 for non-positive integer x+y and not x or y
	{"sc", -1, -4, NaN},          // Beta(-1, -4) = +Inf

	{"sc", +Inf, +Inf, 0},   // Beta(+Inf, +Inf) = 0
	{"sc", +Inf, -Inf, 0},   // Beta(+Inf, -Inf) = 0
	{"sc", -Inf, -Inf, NaN}, // Beta(-Inf, -Inf) = NaN

	{"sc", +Inf, 1.1, 0},     // Beta(+Inf, x > 0) = 0
	{"sc", +Inf, -1.1, +Inf}, // Beta(+Inf, x < 0) = ±Inf for non-integer x
	{"sc", +Inf, -0.1, -Inf}, // Beta(+Inf, x < 0) = ±Inf for non-integer x
	{"sc", +Inf, -2, NaN},    // Beta(+Inf, x ≤ 0) = +Inf for integer x

	{"sc", -Inf, 1.1, 0},    // Beta(-Inf, x > 0) = 0
	{"sc", -Inf, -0.1, NaN}, // Beta(-Inf, x ≤ 0) = NaN

	{"", NaN, 1, NaN},
}

func TestBeta(t *testing.T) { testutil.Test(t, tolBeta, casesBeta, Beta) }

func BenchmarkBeta(b *testing.B) {
	for _, c := range casesBeta {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Beta(c.In1, c.In2)
			}
		})
	}
}
