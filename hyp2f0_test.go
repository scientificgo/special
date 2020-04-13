// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

const tolHyp2F0 = 1 * Macheps

var casesHyp2F0 = []struct {
	Label              string
	In1, In2, In3, Out float64
}{
	{"", 1, 10, 0, 1},
	{"", -9, 10, 10.125, -18778265867784640371.624683},
	{"", 10, -19, -10.125, 1.1375550541875155700872606E+43},
}

func TestHyp2F0(t *testing.T) { testutil.Test(t, tolHyp2F0, casesHyp2F0, Hyp2F0) }

func BenchmarkHyp2F0(b *testing.B) {
	for _, c := range casesHyp2F0 {
		b.Run(c.Label, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Hyp2F0(c.In1, c.In2, c.In3)
			}
		})
	}
}
