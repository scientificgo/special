// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	. "scientificgo.org/special"
	"scientificgo.org/testutils"
	"testing"
)

var casesZeta = []struct {
	Label   string
	In, Out float64
}{
	{"", nan, nan},
	{"", -1, -0.08333333333333333},
	{"", 0, -0.5},
	{"", 1, +inf},
	{"", +inf, 1},
	{"", 1e-08, -0.5000000091893855},
	{"", 0.9999999999, -9.999999172019142e+09},
	{"", 1.0000000001, 9.999999173173574e+09},
	{"", 1.0000001, 1.0000000571377e+07},
	{"", 5, 1.03692775514337},
	{"", 7, 1.008349277381923},
	{"", 9, 1.0020083928260821},
	{"", 10, 1.000994575127818},
	{"", 10.5, 1.000700842641736},
	{"", 11, 1.0004941886041194},
	{"", 12, 1.000246086553308},
	{"", 15, 1.000030588236307},
	{"", 20, 1.0000009539620338},
	{"", 25, 1.0000000298035034},
	{"", 50, 1.0000000000000009},
	{"", 75, 1},
	{"", 750, 1},
	{"", -1.00000001, -0.08333333167912192},
	{"", -2, 0},
	{"", -64, 0},
	{"", -1e+06, 0},
	{"", -1e-08, -0.49999999081061475},
	{"", -1e-06, -0.49999908106247},
	{"", -10.2, 0.004134657457445868},
	{"", -11, 0.021092796092796094},
	{"", -3.2, 0.007011972077091051},
}

var casesEta = []struct {
	Label   string
	In, Out float64
}{
	{"", nan, nan},
	{"", -1, 0.25},
	{"", 0, 0.5},
	{"", 1, 0.6931471805599453},
	{"", +inf, 1},
	{"", -2, 0},
	{"", -1000, 0},
	{"", -1.1, 0.2234614116270079},
	{"", 1.1, 0.7088088499305867},
}

func TestZeta(t *testing.T) { testutils.Test(t, tol, Zeta, casesZeta) }
func TestEta(t *testing.T)  { testutils.Test(t, tol, Eta, casesEta) }

/*
func BenchmarkZeta(b *testing.B) {
	bench(b, cZeta, "",
		func(x []float64) float64 {
			return Zeta(x[0])
		})
}
*/
