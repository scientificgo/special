// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "github.com/scientificgo/special"
)

var casesLgammaRatio = []struct {
	Label    string
	In1, In2 []float64
	Out1     float64
	Out2     int
}{
	{"", []float64{nan}, []float64{0.93}, nan, 1},
	{"", []float64{1}, []float64{nan}, nan, 1},
	{"", []float64{-100}, []float64{-100}, 0, 1},
	{"", []float64{+inf}, []float64{+inf}, 0, 1},
	{"", []float64{-inf}, []float64{-inf}, 0, 1},
	{"", []float64{-10}, []float64{-inf}, nan, 1},
	{"", []float64{17.1}, []float64{+inf}, nan, 1},
	{"", []float64{+inf}, []float64{17.1}, nan, 1},
	{"", []float64{-4}, []float64{-7}, 5.3471075307174685, -1},
	{"", []float64{-100}, []float64{-150}, 241.2807302938602, 1},
	{"", []float64{-3}, []float64{102.3}, +inf, -1},
	{"", []float64{-4}, []float64{102.3}, +inf, 1},
	{"", []float64{102.2132}, []float64{-19}, -inf, -1},
	{"", []float64{18.3}, []float64{17.3}, 2.8507065015037334, 1},
	{"", []float64{18.3}, []float64{19.3}, -2.9069010598473755, 1},
	{"", []float64{17.3}, []float64{18.3}, -2.8507065015037334, 1},
	{"", []float64{3056.675}, []float64{2994.86}, 495.4309926489057, 1},
	{"", []float64{52}, []float64{34}, 67.35512556691585, 1},
	{"", []float64{5200}, []float64{5170}, 256.60281871217694, 1},
	{"", []float64{-3056.675}, []float64{-2994.86}, -496.14583938109644, 1},
	{"", []float64{3056.675}, []float64{-2994.86}, 42451.17912019167, -1},
	{"", []float64{-3056.675}, []float64{2994.86}, -42451.89396692386, -1},
	{"", []float64{305.6675}, []float64{-299.486}, 2852.4000840572594, 1},
	{"", []float64{30.56675}, []float64{-29.9486}, 144.69071788326048, 1},
	{"", []float64{-30.56675}, []float64{29.9486}, -146.5167345931291, -1},
	{"", []float64{-30.56675}, []float64{2.99486}, -76.12202023483091, -1},
	{"", []float64{30.56675}, []float64{-2.99486}, 69.6951395136519, -1},
	{"", []float64{-30.56675}, []float64{-2.99486}, -78.9190498281851, 1},
	{"", []float64{-3.056675}, []float64{-2.99486}, -2.473130942324609, -1},
	{"", []float64{3.056675}, []float64{-2.99486}, -2.739363250496638, -1},
	{"", []float64{-3.056675}, []float64{2.99486}, 0.32389865102957277, 1},
	{"", []float64{-3.056675e-05}, []float64{-2.99486e-05}, -0.02042988114070874, 1},
	{"", []float64{-3.056675e-05}, []float64{2.99486e-05}, -0.020395307538563508, -1},
}

func TestLgammaRatio(t *testing.T) {
	for i, c := range casesLgammaRatio {
		t.Run(c.Label, func(tt *testing.T) {
			res1, res2 := LgammaRatio(c.In1, c.In2)
			ok := equalFloat64(res1, c.Out1) && res2 == c.Out2
			if !ok {
				tt.Errorf("[%v]: Got (%v, %v), want (%v, %v)", i, res1, res2, c.Out1, c.Out2)
			}
		})
	}
}

/*
func BenchmarkLgammaRatio(b *testing.B) {
	GlobalF = bench(b, cLgammaRatio, "",
		func(x []float64) float64 {
			v, _ := LgammaRatio(x[:1], x[1:])
			return v
		})
}
*/
