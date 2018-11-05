// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutils"
)

var casesDigamma = []struct {
	Label   string
	In, Out interface{}
}{
	{"", nan, nan},
	{"", -inf, nan},
	{"", +inf, +inf},
	{"", -2., nan},
	{"", 1e+08, 18.420680738952367},
	{"", 1e-08, -1.0000000057721564e+08},
	{"", -1e-08, 9.999999942278431e+07},
	{"", 10., 2.251752589066721},
	{"", -10.2, 6.6946384504174965},
	{"", 4.9, 1.483737793254897},
	{"", -4.9, -7.981008564556067},
}

func TestDigamma(t *testing.T) { testutils.Test(t, tol, casesDigamma, Digamma) }

/*
func BenchmarkDigamma(b *testing.B) {
	GlobalF = bench(b, cDigamma, "",
		func(x []float64) float64 {
			return Digamma(x[0])
		})
}
*/
