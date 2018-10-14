// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	. "scientificgo.org/special"
	"scientificgo.org/testutils"
	"testing"
)

var casesEi = []struct {
	Label   string
	In, Out float64
}{
	{"", nan, nan},
	{"", 750, +inf},
	{"", +inf, +inf},
	{"", -inf, 0},
	{"", 0, -inf},
	{"", 1e-05, -10.935699800043695},
	{"", 1e-10, -22.448635264938925},
	{"", 0.01, -4.017929465426669},
	{"", 1, 1.8951178163559368},
	{"", 2, 4.95423435600189},
	{"", 5, 40.18527535580318},
	{"", 6, 85.9897621424392},
	{"", 6.01, 86.66495316518093},
	{"", 7.5, 289.3883982001446},
	{"", 8, 440.37989953483816},
	{"", 9, 1037.8782907170896},
	{"", 10, 2492.2289762418777},
	{"", 10.00001, 2492.251002806792},
	{"", 15, 234955.8524907683},
	{"", 20, 2.5615652664056588e+07},
	{"", 25, 3.0059509065255485e+09},
	{"", 27.5, 3.3153303643324104e+10},
	{"", 29.99, 3.654282129023645e+11},
	{"", 30, 3.689732094072742e+11},
	{"", 35, 4.669055014466159e+13},
	{"", 40, 6.039718263611242e+15},
	{"", 45, 7.943916035704454e+17},
	{"", 50, 1.058563689713169e+20},
	{"", 100, 2.71555274485388e+41},
	{"", 200, 3.6312352331593567e+84},
	{"", 705, 2.138148975687854e+303},
	{"", 716, 1.2605029106040893e+308},
	{"", -1e-100, -229.68129363450305},
	{"", -1e-10, -22.448635265138925},
	{"", -1e-05, -10.935719800043696},
	{"", -0.01, -4.037929576538114},
	{"", -1, -0.21938393439552029},
	{"", -2, -0.04890051070806112},
	{"", -4, -0.0037793524098489067},
	{"", -4.0000001, -0.0037793519519579615},
	{"", -5, -0.0011482955912753257},
	{"", -6, -0.0003600824521626587},
	{"", -7, -0.00011548173161033822},
	{"", -8, -3.76656228439249e-05},
	{"", -9, -1.2447354178006272e-05},
	{"", -10, -4.156968929685325e-06},
	{"", -11, -1.4003003042474418e-06},
	{"", -12, -4.7510818246724936e-07},
	{"", -13, -1.6218662188014328e-07},
	{"", -14, -5.5656311111451816e-08},
	{"", -15, -1.918627892147867e-08},
	{"", -20, -9.835525290649882e-11},
	{"", -30, -3.0215520106888124e-15},
	{"", -50, -3.783264029550459e-24},
	{"", -100.5, -2.2232069135926296e-46},
	{"", -105.5, -1.427644740730669e-48},
	{"", -500, -1.4220767822536383e-220},
	{"", -700, -1.406518766234033e-307},
}

func TestEi(t *testing.T) {
	testutils.Test(t, tol, Ei, casesEi)
}

/*
func BenchmarkEi(b *testing.B) {
	GlobalF = bench(b, cEi, "",
		func(x []float64) float64 {
			return Ei(x[0])
		})
}
*/
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
	testutils.Test(t, tol, En, casesEn)
}

/*
func BenchmarkEn(b *testing.B) {
	GlobalF = bench(b, cEn, "",
		func(x []float64) float64 {
			return En(int(x[0]), x[1])
		})
}
*/
var casesLi = []struct {
	Label   string
	In, Out float64
}{
	{"", -1, nan},
	{"", 1e-300, -1.4455586289196509e-303},
	{"", 1e-100, -4.324245377202095e-103},
	{"", 1e-50, -8.611722553713826e-53},
	{"", 1e-25, -1.707997398334023e-27},
	{"", 1e-15, -2.815944392958954e-17},
	{"", 1e-14, -3.0113475954567504e-16},
	{"", 1e-13, -3.2359176090698706e-15},
	{"", 1e-12, -3.4967184067697225e-14},
	{"", 1e-10, -4.168887750019648e-12},
	{"", 1e-08, -5.161659103222967e-10},
	{"", 1e-07, -5.859859508790472e-09},
	{"", 1e-05, -8.036767291447334e-07},
	{"", 2, 1.045163780117493},
	{"", 1e+250, 1.740206254656917e+247},
}

func TestLi(t *testing.T) {
	testutils.Test(t, tol, Li, casesLi)
}

/*
func BenchmarkLi(b *testing.B) {
	GlobalF = bench(b, cLi, "",
		func(x []float64) float64 {
			return Li(x[0])
		})
}
*/

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
	testutils.Test(t, tol, Li2, casesLi2)
}
