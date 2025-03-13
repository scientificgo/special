// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "github.com/scientificgo/special"
)

var casesJacobiP = []struct {
	Label              string
	In1                int
	In2, In3, In4, Out float64
}{
	{"", 2, 1, 1, nan, nan},
	{"", 53, -100, 1, 2.222, nan},
	{"", -2, 1, 1, 2, 0},
	{"", 0, 2.1, 2.3, 3.4, 1},
	{"", 1, 1, 1, -3.141, -6.282},
	{"", 1, 2, 3, 3, 10},
	{"", 2, 0, 10, 3.5, 175.6875},
	{"", 22, -11.5, -51.325, 3.5, 5.02848534357238466003561113176435788066699268664268947e+09},
	{"", 40, 20, 15, 4.32455, 3.7288383613028486979974922944450017018359467994585040e+44},
	{"", 41, 19, 17, -12.12, -1.047080376480031833608451115342239184829088792271637e+64},
	{"", 41, 19, -17, -12.12, -4.266318761115648840302672836461948191620566162114378e+55},
	{"", 41, -17, 19, 12.12, 4.266318761115648840302672836461948191620566162114378e+55},
	{"", 53, 1, -1, 2.222, 1.5050143074960853494568306577510749612555809495544448e+32},
	{"", 53, -1, 1, 2.222, 5.7080306758541784513850001979261750547930475492102161e+31},
	{"", 3, -1, -1, 2.222, 4.374322524},
	{"", 3, -3, 1, 0.222, -0.235455476},
	{"", 3, -3, 1, -0.222, -0.912396524},
	{"", 24, -23, -13, -0.222, 0},
	{"", 24, -10, -13, -0.222, -1.4477482169375526351080034565874610694581906994323984e-07},
}

func TestJacobiP(t *testing.T) {
	for i, c := range casesJacobiP {
		t.Run(c.Label, func(tt *testing.T) {
			res := JacobiP(c.In1, c.In2, c.In3, c.In4)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("[%v]: Got %v, want %v", i, res, c.Out)
			}
		})
	}
}

/*
func BenchmarkJacobiP(b *testing.B) {
	GlobalF = bench(b, cJacobiP, "",
		func(x []float64) float64 {
			return JacobiP(int(x[0]), x[1], x[2], x[3])
		})
}
*/
