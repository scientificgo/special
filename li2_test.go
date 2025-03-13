// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "github.com/scientificgo/special"
)

var casesLi2 = []struct {
	Label   string
	In, Out float64
}{
	{"", -1, nan},
	{"", +inf, +inf},
	{"", 1, -inf},
	{"", 1e-321, -1.045163780117493},
	{"", 1e-100, -1.045163780117493},
	{"", 1e-10, -1.0451637801216618},
	{"", 1e+250, 1.740206254656917e+247},
}

func TestLi2(t *testing.T) {
	for i, c := range casesLi2 {
		t.Run(c.Label, func(tt *testing.T) {
			res := Li2(c.In)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("[%v]: Got %v, want %v", i, res, c.Out)
			}
		})
	}
}
