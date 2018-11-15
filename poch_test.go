// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

var casesPoch = []struct {
	Label         string
	In1, In2, Out float64
}{
	{"", nan, 45.456789, nan},
	{"", 45.456789, nan, nan},
	{"", 45.456789, 0, 1},
	{"", 45.456789, 1, 45.456789},
	{"", 1, -1, +inf},
	{"", -7, +inf, 0},
	{"", -6.99, +inf, -inf},
	{"", 45.456789, +inf, +inf},
	{"", 45.456789, -inf, nan},
	{"", +inf, 45.456789, +inf},
	{"", -inf, 45.456789, 0},
	{"", +inf, +inf, +inf},
	{"", 3, -11, +inf},
	{"", 3, -10, -inf},
	{"", -1.25, -2.75, +inf},
	{"", 1.25, -2.25, -inf},
	{"", -0.1, -0.9, +inf},
	{"", 0, -5, -0.008333333333333333},
	{"", 0, -4.9, 0},
	{"", -5, 5, -120},
	{"", 2, 5, 720},
	{"", -3, 8, 0},
	{"", -11, 45.456789, 0},
	{"", 8, -3, 0.004761904761904762},
	{"", -10, 4, 5040},
	{"", -5, -10, 9.17659647818378e-11},
	{"", 1.25, 2.75, 6.619575907925023},
	{"", -17.5, 180.5, 5.86592464809013e+303},
	{"", 6, 45.456789, 1.5234777676342018e+63},
	{"", 140, -190.1, -2.348215044179754e-303},
	{"", -3.141592653589793, 1.5707963267948966, 2.266906744884818},
	{"", 3.141592653589793, 3.141592653589793, 85.63499961560903},
	{"", 200.1, -0.1, 0.5888365802348768},
	{"", 200, -0.1, 0.5888660965321199},
	{"", -205.1, 131.9, 2.305030399715578e+280},
	{"", -0.1, 0.2, -0.8902538065654994},
	{"", -1e-50, 190, -5.0949066711869736e+299},
}

func TestPoch(t *testing.T) {
	testutil.Test(t, tol, casesPoch, Poch)
}
