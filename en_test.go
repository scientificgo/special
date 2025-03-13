// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "github.com/scientificgo/special"
)

var casesEn = []struct {
	Label    string
	In1      int
	In2, Out float64
}{
	{"", 1, nan, nan},
	{"", -1, 2, nan},
	{"", 1, +inf, 0},
	{"", 1, -1, -1.8951178163559368},
	{"", 1, 0, +inf},
	{"", 1, -720, -inf},
	{"", 2, -1, nan},
	{"", 1, -100000, -inf},
	{"", 5, 0, 0.25},
	{"", 0, 1e-08, 9.9999999e+07},
	{"", 2, 1e-08, 0.9999998115653491},
	{"", 3, 1e-08, 0.49999999000000095},
	{"", 4, 1e-08, 0.3333333283333334},
	{"", 5, 1e-08, 0.2499999966666667},
	{"", 6, 1e-08, 0.19999999750000003},
	{"", 7, 1e-08, 0.1666666646666667},
	{"", 8, 1e-08, 0.1428571411904762},
	{"", 9, 1e-08, 0.12499999857142859},
	{"", 10, 1e-08, 0.11111110986111111},
	{"", 20, 1e-08, 0.052631578391812865},
	{"", 100, 1e-08, 0.010101009998969285},
	{"", 2, 1, 0.14849550677592205},
	{"", 5, 1, 0.0704542374617204},
	{"", 10, 1, 0.0363939940314164},
	{"", 20, 1, 0.018345971206755872},
	{"", 100, 1, 0.003678422930396813},
	{"", 2, 1.99999, 0.03753475082898094},
	{"", 2, 2, 0.03753426182049045},
	{"", 5, 2, 0.02132240020232302},
	{"", 10, 2, 0.012092085136400296},
	{"", 20, 2, 0.0064143058553249},
	{"", 100, 2, 0.0013396881216721757},
	{"", 2, 5, 0.000996469042708838},
	{"", 5, 5, 0.0007057606934245853},
	{"", 10, 5, 0.0004691048076578119},
	{"", 20, 5, 0.0002782745928857308},
	{"", 100, 5, 6.475775388266509e-05},
	{"", 2, 25, 5.156944766134789e-13},
	{"", 3, 25, 4.977909748135229e-13},
	{"", 4, 25, 4.810564982086496e-13},
	{"", 5, 25, 4.653828524369453e-13},
	{"", 6, 25, 4.5067451080807775e-13},
	{"", 7, 25, 4.3684684912701285e-13},
	{"", 8, 25, 4.238246623983856e-13},
	{"", 9, 25, 4.115409131255475e-13},
	{"", 10, 25, 3.999356707583703e-13},
	{"", 11, 25, 3.8895520960047644e-13},
	{"", 20, 25, 3.1162408463648225e-13},
	{"", 30, 25, 2.5499537807523303e-13},
	{"", 50, 25, 1.8681837014276917e-13},
	{"", 100, 25, 1.1181687061260001e-13},
	{"", 2, 500, 1.4192495473093422e-220},
	{"", 5, 500, 1.4108347621366478e-220},
	{"", 10, 500, 1.3970292719286189e-220},
	{"", 20, 500, 1.3702118216839137e-220},
	{"", 1000, 500, 4.75182858947099e-221},
}

func TestEn(t *testing.T) {
	for i, c := range casesEn {
		t.Run(c.Label, func(tt *testing.T) {
			res := En(c.In1, c.In2)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("[%v]: Got %v, want %v", i, res, c.Out)
			}
		})
	}
}

/*
func BenchmarkEn(b *testing.B) {
	GlobalF = bench(b, cEn, "",
		func(x []float64) float64 {
			return En(int(x[0]), x[1])
		})
}
*/
