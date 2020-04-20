// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

// Rearrange Stirling's approximation to get:
//
//  Stirling(x) ~ Gamma(x) * Sqrt(x) * (e/x)**x
//

const tolStirling = 1 * Macheps

var casesStirling = []struct {
	Label   string
	In, Out float64
}{
	{"", 100, 2.508717995156920341603935},
	{"", 150, 2.5080212307274908952840762},
	{"", 160, 2.5079341485321726572651406},
	{"", 170.25, 2.5078555085801944460726609},
	{"", 190.25, 2.507726467807611103090845},
	{"", 200.06125, 2.5076725999368071608958059},
	{"", 5000.06125, 2.5066700516052278677031586},
	{"", 1000000.5, 2.506628483516594315762},
	{"", 13640500.5, 2.50662828994463875129},
}

func TestStirling(t *testing.T) { testutil.Test(t, tolStirling, casesStirling, Stirling) }
