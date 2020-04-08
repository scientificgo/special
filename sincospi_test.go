// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

const tolSinPi = 1 * Macheps

var casesSinPi = []struct {
	Label   string
	In, Out float64
}{
	{"", 0.25, 0.70710678118654752440084436},
	{"", 1.875, -0.38268343236508977172845998},
	{"", 84.7390380859375, 0.73103395751020752587782301},

	// special cases
	{"sc", 0, 0},
	{"sc", NegativeZero, NegativeZero},
	{"sc", 1, 0},
	{"sc", 1e10, 0},
	{"sc", -55.000000000000001, 0},
	{"sc", 101.5, -1},
	{"sc", +Inf, NaN},
	{"sc", -Inf, NaN},
	{"sc", NaN, NaN},
}

func TestSinPi(t *testing.T) { testutil.Test(t, tolSinPi, casesSinPi, SinPi) }
func BenchmarkSinPi(b *testing.B) {
	for _, c := range casesSinPi {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = SinPi(c.In)
			}
		})
	}
}

const tolCosPi = 1 * Macheps

var casesCosPi = []struct {
	Label   string
	In, Out float64
}{
	{"", 0.25, 0.70710678118654752440084436},
	{"", 1.875, 0.92387953251128675612818319},
	{"", 84.7390380859375, -0.68234108257305165552380618},

	// special cases
	{"sc", 0, 1},
	{"sc", NegativeZero, 1},
	{"sc", 1, -1},
	{"sc", 1e10, 1},
	{"sc", -55.000000000000001, -1},
	{"sc", 101.5, 0},
	{"sc", +Inf, NaN},
	{"sc", -Inf, NaN},
	{"sc", NaN, NaN},
}

func TestCosPi(t *testing.T) { testutil.Test(t, tolCosPi, casesCosPi, CosPi) }

func BenchmarkCosPi(b *testing.B) {
	for _, c := range casesCosPi {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = CosPi(c.In)
			}
		})
	}
}

func TestSincosPi(t *testing.T) {
	sinPi := func(x float64) float64 {
		sin, _ := SincosPi(x)
		return sin
	}
	cosPi := func(x float64) float64 {
		_, cos := SincosPi(x)
		return cos
	}
	testutil.Test(t, tolSinPi, casesSinPi, sinPi)
	testutil.Test(t, tolCosPi, casesCosPi, cosPi)
}

const tolTanPi = 1 * Macheps

var casesTanPi = []struct {
	Label   string
	In, Out float64
}{
	{"", 0.25, 1},
	{"", 1.875, -0.41421356237309504880168872},
	{"", 84.7390380859375, -1.0713614879431545813733519},

	// special cases
	{"sc", 0, 0},
	{"sc", NegativeZero, NegativeZero},
	{"sc", 1, 0},
	{"sc", 1e10, 0},
	{"sc", -55.000000000000001, 0},
	{"sc", 101.5, +Inf},
	{"sc", +Inf, NaN},
	{"sc", -Inf, NaN},
	{"sc", NaN, NaN},
}

func TestTanPi(t *testing.T) { testutil.Test(t, tolTanPi, casesTanPi, TanPi) }

func BenchmarkTanPi(b *testing.B) {
	for _, c := range casesTanPi {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = TanPi(c.In)
			}
		})
	}
}

const tolCotPi = 1 * Macheps

var casesCotPi = []struct {
	Label   string
	In, Out float64
}{
	// Laurent series
	{"", 0.499999999999, 3.1415231561563747518609626E-12},
	{"", 0.49609645666800606, 0.012263957851427457500426728},
	{"", 0.4953578784958714, 0.014584688800832377749780422},
	{"", 0.49447955607858307, 0.017344725080262040179033983},
	{"", 0.45465857480079436, 0.14341558830980207152883813},
	{"", 0.40379574607944585, 0.31178634892876441389629369},

	// ratio of cos/sin
	{"", 0.3446558574800795, 0.53085778685556093753773502},
	{"", 0.30749447955607856, 0.69116815987925748761345253},
	{"", 0.2619344658939359, 0.92769076033210862733002030},
	{"", 0.2045746199463584, 1.3355885103240076231619590},
	{"", 0.15303002991764028, 1.9172709386974155603115287},
	{"", 0.11784925006577396, 2.5764379030830447323146581},
	{"", 0.04028720855712891, 7.8587823787168608866418159},
	{"", 7.888609052210118e-31, 4.0350571827946162986346419E+29},
	{"", -9.999778782798783e-13, -3.18316927901779651687601719E+11},

	// standard cases
	{"", 0.25, 1},
	{"", 1.875, -2.4142135623730950488016887},
	{"", 84.7390380859375, -0.93339177416190551563386356},

	// special cases
	{"sc", 0, +Inf},
	{"sc", NegativeZero, -Inf},
	{"sc", 1, +Inf},
	{"sc", 1e10, +Inf},
	{"sc", -55.000000000000001, +Inf},
	{"sc", 101.5, 0},
	{"sc", +Inf, NaN},
	{"sc", -Inf, NaN},
	{"sc", NaN, NaN},
}

func TestCotPi(t *testing.T) { testutil.Test(t, tolCotPi, casesCotPi, CotPi) }

func BenchmarkCotPi(b *testing.B) {
	for _, c := range casesCotPi {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = CotPi(c.In)
			}
		})
	}
}
