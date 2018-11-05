// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutils"
)

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

func TestEta(t *testing.T) { testutils.Test(t, tol, casesEta, Eta) }
