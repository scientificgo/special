// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

var casesFibonacci = []struct {
	Label   string
	In, Out float64
}{
	{"", +inf, +inf},
	{"", -inf, nan},
	{"", nan, nan},
	{"", 0, 0},
	{"", 1, 1},
	{"", 2, 1},
	{"", -1, 1},
	{"", -2, -1},
	{"", 2.5, 1.4893065462657091},
	{"", -123.321, 1.4123645428734893e+25},
	{"", -100, -3.542248481792619e+20},
	{"", 1.618033988749895, 0.8998442262232443},
	{"", 87, 6.798916376386122e+17},
}

func TestFibonacci(t *testing.T) {
	testutil.Test(t, tol, casesFibonacci, Fibonacci)
}
