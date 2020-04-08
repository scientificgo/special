// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

const tolTrigamma = 2 * Macheps

var casesTrigamma = []struct {
	Label   string
	In, Out float64
}{
	{"", 1.463012695e-08, 4.672010506017306919879455074747733792395492208426383879009e+15},
	{"", -1.463012695e-08, 4.672010506017306919879525419728111187299900432232048543168e+15},
	{"", 1., 1.644934066848226436472415166646025189218949901206798437736e+0},
	{"", -1.0000000001, 9.99999834519278539750415758026360397410756315570886797784e+19},
	{"", 1.000001, 1.64493166273766728061509770051451795896535999183818024319e+0},
	{"", -1.000001, 1.0000000001671782103838804778917733132247597234668371e+12},
	{"", 10.2, 1.030018134500800395758617117236581251106976298063732535426e-1},
	{"", -10.2, 2.847346121728795471557735494983233945229243383837781813483e+1},
	{"", 1.000000001e+06, 1.000000499000165620168669424464920305836499902399596165282e-6},
	{"", -1.000000001e+06, 1.000003194878731804390881683176548086346905952327605006496e+6},
	{"", 9.1773313224759e+08, 1.089641384196028394834067488737338031486320209571485363442e-9},
	{"", -9.1773313224759e+08, 2.004270014626729594326872621144420283846507969108883303604e+1},

	// Special cases
	{"sc", 0., +Inf},   // Trigamma(-k) = +Inf for integer k ≥ 0
	{"sc", -11., +Inf}, // Trigamma(-k) = +Inf for integer k ≥ 0
	{"sc", +Inf, 0.},   // Trigamma(+Inf) = 0
	{"sc", -Inf, NaN},  // Trigamma(-Inf) = NaN
	{"sc", NaN, NaN},
}

func TestTrigamma(t *testing.T) { testutil.Test(t, tolTrigamma, casesTrigamma, Trigamma) }
func BenchmarkTrigamma(b *testing.B) {
	for _, c := range casesTrigamma {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Trigamma(c.In)
			}
		})
	}
}
