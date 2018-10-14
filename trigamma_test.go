// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	. "scientificgo.org/special"
	"scientificgo.org/testutils"
	"testing"
)

var casesTrigamma = []struct {
	Label   string
	In, Out interface{}
}{
	{"", nan, nan},
	{"", -inf, nan},
	{"", +inf, 0.},
	{"", -2., nan},
	{"", 1e+08, 1.000000005e-08},
	{"", 1e-08, 1.0000000000000002e+16},
	{"", -1e-08, 1.0000000000000002e+16},
	{"", 10., 0.10516633568168575},
	{"", -10.2, 28.473461217287777},
	{"", 4.9, 0.22631146419168002},
	{"", -4.9, 103.17117695886462},
}

func TestTrigamma(t *testing.T) {
	testutils.Test(t, tol, Trigamma, casesTrigamma)
}

/*
func BenchmarkTrigamma(b *testing.B) {
	GlobalF = bench(b, cTrigamma, "",
		func(x []float64) float64 {
			return Trigamma(x[0])
		})
}
*/
