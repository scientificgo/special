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

const tolEllipticK = 2 * Macheps

var casesEllipticK = []struct {
	Label   string
	In, Out float64
}{
	// Print inputs with "%.50f" before using them to verify results
	// as the function is _extremely_ sensitive near x ~ 1.
	{"", 1e-5, 1.5708002538078030832179975},
	{"", 0.1, 1.6124413487202194006835223},
	{"", 0.2, 1.6596235986105280064487208},
	{"", 0.5, 1.8540746773013719184338504},
	{"", 0.9, 2.57809211334817329268264},
	{"", 0.999, 4.8411325605502965869924175},
	{"", 0.99999, 7.1427724505840536390265965},
	{"", 0.9999999, 9.4453423979957946391774502},
	{"", 0.999999999, 11.747927296421043877523451},
	{"", 0.99999999999, 14.050512331249584376408119},
	{"", 0.99999999999999, 17.50478981079335001398162},
	{"", 0.9999999999999999, 19.754694645958441838938461},

	{"", -1e-5, 1.5707923998261688019381013},
	{"", -0.1, 1.5335928197134568795403596},
	{"", -0.2, 1.5000268912867475170203126},
	{"", -0.5, 1.415737208425956198892166},
	{"", -0.9, 1.3293621856564093582898001},
	{"", -1, 1.3110287771460599052324198},
	{"", -10, 0.79087189023873847519891547},
	{"", -100, 0.36821924860914103291985717},
	{"", -1e3, 0.15302933498849878576726703},
	{"", -1e4, 0.059913397672787116807457818},
	{"", -1e6, 0.0082940478165906199329226377},
	{"", -1e10, 1.2899219825792638543288667e-4},
	{"", -1e20, 2.4412145291060347458955848e-9},
	{"", -1e100, 1.1651554901082217481973404e-48},
	{"", -1e200, 2.3164480366052445902063361e-98},
	{"", -1e300, 3.4677405831022674322153318e-148},
	{"", -1e305, 1.1147993912208864662153783e-150},
	{"", -1.7976931348623157e+308, 2.6572401146362278002845198e-152},

	// special cases
	{"sc", 0, math.Pi / 2}, // K(0) = π / 2
	{"sc", 1, +Inf},        // K(1) = +Inf
	{"", -Inf, 0},          // K(-Inf) = 0
	{"", +Inf, NaN},        // K(x > 1) = NaN
	{"sc", 1.2, NaN},       // K(x > 1) = NaN
	{"sc", NaN, NaN},
}

func TestEllipticK(t *testing.T) { testutil.Test(t, tolEllipticK, casesEllipticK, EllipticK) }

// func BenchmarkEllipticK(b *testing.B) { testutil.Benchmark(b, casesEllipticK, EllipticK) }
func BenchmarkEllipticK(b *testing.B) {
	for _, c := range casesEllipticK {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = EllipticK(c.In)
			}
		})
	}
}
