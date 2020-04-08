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

const tolGCF = 2 * Macheps

var _ones = func(int) float64 { return 1 }
var _twos = func(int) float64 { return 2 }
var _odds = func(i int) float64 { return float64(2*i - 1) }
var _squares = func(i int) float64 {
	i--
	return float64(i) * float64(i)
}

var casesGCF = []struct {
	Label    string
	In1, In2 func(int) float64
	Out      float64
}{
	{"", _ones, _ones, math.Phi - 1},
	{"", _ones, _twos, math.Sqrt2 - 1},
	{"", _squares, _odds, math.Pi / 4},
	{"", _ones, func(int) float64 { return NaN }, NaN},
}

func TestGCF(t *testing.T) { testutil.Test(t, tolGCF, casesGCF, GCF) }

// func BenchmarkGCF(b *testing.B) { testutil.Benchmark(b, casesGCF, GCF) }
func BenchmarkGCF(b *testing.B) {
	for _, c := range casesGCF {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = GCF(c.In1, c.In2)
			}
		})
	}
}
