// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

var casesSi = []struct {
	Label   string
	In, Out float64
}{
	{"", nan, nan},
	{"", -inf, -1.5707963267948966},
	{"", 0, 0},
	{"", 1, 0.946083070367183},
	{"", 2, 1.6054129768026948},
	{"", 4, 1.7582031389490531},
	{"", 5, 1.549931244944674},
	{"", 10, 1.6583475942188741},
	{"", 15, 1.6181944437083688},
	{"", 20, 1.54824170104344},
	{"", 25, 1.5314825509999612},
	{"", 1000, 1.5702331219687713},
	{"", -4.0001, -1.7581842183061578},
	{"", -25, -1.5314825509999612},
}

func TestSi(t *testing.T) {
	testutil.Test(t, tol, casesSi, Si)
}

/*
func BenchmarkSi(b *testing.B) {
	bench(b, cSi, "",
		func(x []float64) float64 {
			return Si(x[0])
		})
}
*/

var casesCi = []struct {
	Label   string
	In, Out float64
}{
	{"", -1, nan},
	{"", +inf, 0},
	{"", 1, 0.33740392290096816},
	{"", 2, 0.422980828774865},
	{"", 4.0001, -0.14099803782717715},
	{"", 1000, 0.0008263155110906822},
}

func TestCi(t *testing.T) {
	testutil.Test(t, tol, casesCi, Ci)
}

/*
func BenchmarkCi(b *testing.B) {
	bench(b, cCi, "",
		func(x []float64) float64 {
			return Ci(x[0])
		})
}
*/

var casesCin = []struct {
	Label   string
	In, Out float64
}{
	{"", 0, 0},
	{"", 4, 2.104491723908354},
	{"", -4, 2.104491723908354},
	{"", 10, 2.9252571909000338},
	{"", -10, 2.9252571909000338},
}

func TestCin(t *testing.T) {
	testutil.Test(t, tol, casesCin, Cin)
}

/*
func BenchmarkCin(b *testing.B) {
	bench(b, cCin, "",
		func(x []float64) float64 {
			return Cin(x[0])
		})
}
*/
