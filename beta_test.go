// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

var casesBeta = []struct {
	Label         string
	In1, In2, Out float64
}{
	{"", nan, 1, nan},
	{"", -inf, -inf, nan},
	{"", +inf, +inf, 0},
	{"", +inf, 190.7, 0},
	{"", 0.986, -inf, nan},
	{"", -3, -7, +inf},
	{"", 10, -10, 0.1},
	{"", -15, 15, -0.06666666666666667},
	{"", 1.5, -2.5, 0},
	{"", 3.54, -1, -inf},
	{"", -1, 3.54, -inf},
	{"", 0, 12, +inf},
	{"", +inf, -7, -inf},
	{"", -8, +inf, +inf},
	{"", 8, +inf, 0},
	{"", 300, 200, 1.6485491608664747e-147},
	{"", 0.5, -0.75, 1.74803836952808},
	{"", 100000, 1e-05, 99987.91060292121},
	{"", -1.123, -1.132, -33.830986471614295},
	{"", -112.3, -113.2, -2.668986182849379e+67},
}

func TestBeta(t *testing.T) {
	testutil.Test(t, tol, casesBeta, Beta)
}
