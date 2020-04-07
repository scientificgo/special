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

var tolSi = 2 * macheps

var casesSi = []struct {
	Label   string
	In, Out float64
}{
	{"", 8972.573293986234, 1.5706867688867213444469011},
	{"", 27.120982110500336, 1.5844721955887463991091255},
	{"", 6.978869287425517, 1.4526346269923779640373967},
	{"", 4.000000000000001, 1.7582031389490528900616482},
	{"", 3.99999999999999, 1.7582031389490549906105367},
	{"", 2.0140909664847597, 1.6117761827221895023155436},
	{"", 0.9989121761095486, 0.94516751998249639763780650},
	{"", 0.5037924372281416, 0.49674262890805088989323438},
	{"", 0.000346453504773739, 0.00034645350246347097152844796},
	{"", 5.527147875260445e-76, 5.5271478752604445602472652E-76},

	// Special cases
	{"sc", 0, 0},              // Si(0) = 0
	{"sc", +inf, math.Pi / 2}, // Si(±inf) = —π/2
	{"sc", nan, nan},
}

func TestSi(t *testing.T) { testutil.Test(t, tolSi, casesSi, Si) }
func BenchmarkSi(b *testing.B) {
	for _, c := range casesSi {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Si(c.In)
			}
		})
	}
}

var tolCi = 2 * macheps

var casesCi = []struct {
	Label   string
	In, Out float64
}{
	{"", 8972.573293986234, 0.000020453200152228253734021304},
	{"", 27.120982110500336, 0.034162162380982727369646079},
	{"", 6.978869287425517, 0.074395226828159826750107085},
	{"", 4.000000000000001, -0.14098169788693055677718446},
	{"", 3.99999999999999, -0.14098169788692874255169000},
	{"", 2.0140909664847597, 0.42001432516058014276199547},
	{"", 0.9989121761095486, 0.33681535109906093669466307},
	{"", 0.5037924372281416, -0.17115971869087344110272626},
	{"", 0.000346453504773739, -7.3905462988011259658505904},
	{"", 5.527147875260445e-76, -172.70957947508479449370152},

	// Special cases
	{"sc", 0, -inf},    // Ci(0) = -inf
	{"sc", +inf, 0},    // Ci(+inf) = 0
	{"sc", -1.23, nan}, // Ci(x) = NaN for x < 0
	{"sc", nan, nan},
}

func TestCi(t *testing.T) { testutil.Test(t, tolCi, casesCi, Ci) }
func BenchmarkCi(b *testing.B) {
	for _, c := range casesCi {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Ci(c.In)
			}
		})
	}
}

var tolCin = 1 * macheps

var casesCin = []struct {
	Label   string
	In, Out float64
}{
	{"", 1e305, 702.86566902808546648609391},
	{"", 1e200, 461.09423426371066966420480},
	{"", 8972.573293986234, 9.6791230034260547167781246},
	{"", 4.000000000000001, 2.1044917239083542582627657},
	{"", 3.99999999999999, 2.1044917239083496684797097},
	{"", 0.5037924372281416, 0.062784456959459733575088805},
	{"", 5.527147875260445e-76, 0.},

	// Special cases
	{"sc", 0, 0},       // Cin(0) = 0
	{"sc", -inf, +inf}, // Cin(±inf) = +inf
	{"sc", nan, nan},
}

func TestCin(t *testing.T) { testutil.Test(t, tolCin, casesCin, Cin) }
func BenchmarkCin(b *testing.B) {
	for _, c := range casesCin {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Cin(c.In)
			}
		})
	}
}
