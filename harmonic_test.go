// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

var casesHarmonic = []struct {
	Label   string
	In, Out float64
}{
	{"", -inf, nan},
	{"", 0, 0},
	{"", 1, 1},
	{"", 2, 1.5},
	{"", 20, 3.597739657143682},
	{"", 50, 4.499205338329425},
}

func TestHarmonic(t *testing.T) { testutil.Test(t, tol, casesHarmonic, Harmonic) }
