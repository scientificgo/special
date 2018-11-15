// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

var casesLegendreQ = []struct {
	Label    string
	In1      int
	In2, Out float64
}{
	{"", 2, nan, nan},
	{"", -2, 2, nan},
	{"", 1, 1.1, nan},
	{"", 0, 0.9, 1.472219489583220230004513715943926768618689630649564409268},
	{"", 1, 0.999, 2.796400966082949831744191300541195457801412019535269391030},
	{"", 11, 0.999, 0.665248555792627905833229643332143091673876114974304501168},
	{"", 101, -0.10101, 0.082745695703743357501272084016026991789721245950554478861},
}

func TestLegendreQ(t *testing.T) { testutil.Test(t, tol, casesLegendreQ, LegendreQ) }

/*
func BenchmarkLegendreQ(b *testing.B) {
	GlobalF = bench(b, cLegendreQ, "",
		func(x []float64) float64 {
			return LegendreQ(int(x[0]), x[1])
		})
}
*/
