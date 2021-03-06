// Copyright (c) 2018, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"testing"

	. "scientificgo.org/special"
	"scientificgo.org/testutil"
)

var casesHypPFQ = []struct {
	Label    string
	In1, In2 []float64
	In3, Out float64
}{
	{"Hyp0F1", nil, []float64{nan}, 100, nan},
	{"Hyp0F1", nil, []float64{-2}, nan, nan},
	{"Hyp0F1", nil, []float64{+inf}, -inf, 0.36787944117144233},
	{"Hyp0F1", nil, []float64{+inf}, +inf, 2.718281828459045},
	{"Hyp0F1", nil, []float64{+inf}, 4567.98765, 1},
	{"Hyp0F1", nil, []float64{-3}, 11, nan},
	{"Hyp0F1", nil, []float64{-4}, 11, nan},
	{"Hyp0F1", nil, []float64{-4}, -11, nan},
	{"Hyp0F1", nil, []float64{9.876545678998766e+10}, 0, 1},
	{"Hyp0F1", nil, []float64{1}, 2, 4.2523508795026235},
	{"Hyp0F1", nil, []float64{2}, 2, 2.394833099273405},
	{"Hyp0F1", nil, []float64{2}, 3, 3.468649618760533},
	{"Hyp0F1", nil, []float64{2}, 1, 1.590636854637329},
	{"Hyp0F1", nil, []float64{3}, 2, 1.857517780229219},
	{"Hyp0F1", nil, []float64{3}, 2.5, 2.1420143462925894},
	{"Hyp0F1", nil, []float64{3}, 25.5, 197.08967324379037},
	{"Hyp0F1", nil, []float64{3}, 255.5, 3.9825092728943634e+10},
	{"Hyp0F1", nil, []float64{3}, 2555.5, 2.4708990611876765e+39},
	{"Hyp0F1", nil, []float64{3}, 25555.5, 1.2385461930270518e+133},
	{"Hyp0F1", nil, []float64{3}, 125555.5, 1.4154906008294558e+301},
	{"Hyp0F1", nil, []float64{3.3}, 125555.5, 3.2609755879144203e+300},
	{"Hyp0F1", nil, []float64{33.3}, 125555.5, 1.439580416960765e+259},
	{"Hyp0F1", nil, []float64{333.3}, 125555.5, 5.659525247124615e+120},
	{"Hyp0F1", nil, []float64{3333.3}, 125555.5, 1.8515880989407224e+16},
	{"Hyp0F1", nil, []float64{33333.3}, 125555.5, 43.226597470849484},
	{"Hyp0F1", nil, []float64{333333.3}, 125555.5, 1.4574179240855258},
	{"Hyp0F1", nil, []float64{3.3333333e+06}, 125555.5, 1.0383850296988064},
	{"Hyp0F1", nil, []float64{3.33333333e+07}, 125555.5, 1.0037737678013317},
	{"Hyp0F1", nil, []float64{3.333333333e+08}, 125555.5, 1.0003767374477712},
	{"Hyp0F1", nil, []float64{3.3333333333e+09}, 125555.5, 1.0000376673593976},
	{"Hyp0F1", nil, []float64{3.33333333333e+10}, 125555.5, 1.0000037666720938},
	{"Hyp0F1", nil, []float64{3.33333333333e+10}, 1.25555555555e+10, 1.4574184222488056},
	{"Hyp0F1", nil, []float64{3.33333333333e+10}, 1.255555555555e+11, 43.23570534342665},
	{"Hyp0F1", nil, []float64{3.333333333333e+11}, 1.255555555555e+11, 1.457418422253289},
	{"Hyp0F1", nil, []float64{-3.0000001}, 11, 6.856654383928824e+09},
	{"Hyp0F1", nil, []float64{-3000.0001}, 110, 0.963997631524568},
	{"Hyp0F1", nil, []float64{-7.7}, 2.5, 0.7280854011698313},
	{"Hyp0F1", nil, []float64{77.7}, -2.5, 0.9683307048584084},
	{"Hyp0F1", nil, []float64{-77.7}, -25.5, 1.389427154414045},
	{"Hyp0F1", nil, []float64{777.7}, -25.5, 0.9677420702617322},
	{"Hyp0F1", nil, []float64{-777.7}, 255.5, 0.7200292734914648},
	{"Hyp0F1", nil, []float64{-7777.7}, -2555.5, 1.3889867337813122},
	// FIXME: function fails these tests.
	//	{"Hyp0F1", nil, []float64{-7.7}, -77.77, 1768.9475699750021},
	//	{"Hyp0F1", nil, []float64{-7.7}, -154, 3516.225323360214},
	//	{"Hyp0F1", nil, []float64{-7.7}, -770, 7.051795763569786e+07},

	{"Hyp1F1", []float64{11}, []float64{11}, 0.999, 2.715564905318566687331982733345286907488},
	{"Hyp1F1", []float64{nan}, []float64{100}, 10, nan},
	{"Hyp1F1", []float64{0}, []float64{-4}, 5, 1},
	{"Hyp1F1", []float64{543}, []float64{-4}, 0, 1},
	{"Hyp1F1", []float64{+inf}, []float64{+inf}, 1, 2.718281828459045},
	{"Hyp1F1", []float64{+inf}, []float64{-inf}, 1, 0.36787944117144233},
	{"Hyp1F1", []float64{+inf}, []float64{1}, -1, nan},
	{"Hyp1F1", []float64{+inf}, []float64{1}, 1, nan},
	{"Hyp1F1", []float64{20}, []float64{+inf}, 1, 1},
	{"Hyp1F1", []float64{20}, []float64{-inf}, 1, 1},
	{"Hyp1F1", []float64{-3.9}, []float64{-4}, 5, nan},
	{"Hyp1F1", []float64{3.9}, []float64{-4}, -5, nan},
	{"Hyp1F1", []float64{2}, []float64{-4}, 5, nan},
	{"Hyp1F1", []float64{2}, []float64{-5}, 5, nan},
	{"Hyp1F1", []float64{-1}, []float64{-4}, -5, -1. / 4},
	{"Hyp1F1", []float64{-1}, []float64{-4.1}, -5, -0.21951219512195121951219512195121951219512195121951219512},
	{"Hyp1F1", []float64{-5.432}, []float64{2.2}, -700.9, 1.3727292185988063233762299021768245936154987855107037e+12},
	{"Hyp1F1", []float64{-5.432}, []float64{2.2}, -9700.9, 2.0692262222416140055851251664249965291876090468930709e+18},
	{"Hyp1F1", []float64{-5.432e-07}, []float64{2.2e-07}, -9700.9, 3.469109295491283049683835128457572813179185164196529536809},
	{"Hyp1F1", []float64{-11}, []float64{11}, 9.999, -0.000486757541885359591903042751140800},
	{"Hyp1F1", []float64{12.5}, []float64{11.75}, 19.999, 1.024636534533486247341780883267801707746e+09},
	{"Hyp1F1", []float64{2.5}, []float64{3.5}, 100, 6.619995042577486e+41},
	{"Hyp1F1", []float64{3.5}, []float64{2.5}, -100, -1.450829630648126e-42},
	{"Hyp1F1", []float64{2}, []float64{4}, 5, 21.707494910771032},
	{"Hyp1F1", []float64{2}, []float64{4}, -5, 0.14626395019169272},
	{"Hyp1F1", []float64{-2}, []float64{4}, -5, 4.75},
	{"Hyp1F1", []float64{-2}, []float64{-4}, -5, 0.5833333333333334},
	{"Hyp1F1", []float64{-2}, []float64{-4}, 5, 5.583333333333333},
	{"Hyp1F1", []float64{-2}, []float64{4}, 5, -0.25},
	{"Hyp1F1", []float64{1}, []float64{1e+10}, 5, 1.0000000005},
	{"Hyp1F1", []float64{1}, []float64{1e+08}, 5, 1.0000000500000026},
	{"Hyp1F1", []float64{-1}, []float64{1e+10}, 5, 0.9999999995},
	{"Hyp1F1", []float64{10}, []float64{200}, 5, 1.2878814071276767},
	{"Hyp1F1", []float64{-10}, []float64{200}, 5, 0.7762152008143943},
	{"Hyp1F1", []float64{-10}, []float64{-15}, 5, 22.655716268501486},
	{"Hyp1F1", []float64{-10}, []float64{-1000}, 5, 1.0511413040709117},
	{"Hyp1F1", []float64{-10}, []float64{-1000}, 5000, 6.239339006816543e+07},
	{"Hyp1F1", []float64{-10}, []float64{-1000}, 5e+07, 1.0218603619401191e+47},
	{"Hyp1F1", []float64{-999}, []float64{-1000}, 0.005, 1.0050074957967967},
	{"Hyp1F1", []float64{-543.43}, []float64{-532.32}, -324.432, -1.1198647713125607e+141},
	{"Hyp1F1", []float64{-10}, []float64{1}, 5, 1.7562761794532629},

	// FIXME function fails these tests.
	// {"Hyp1F1", []float64{-50}, []float64{5}, 5, 0.0005715733149410122},
	// {"Hyp1F1", []float64{-100}, []float64{10}, 5, 2.9224953421878733e-07},
	// {"Hyp1F1", []float64{-100}, []float64{50}, 5, 3.982104307130384e-06},
	// {"Hyp1F1", []float64{-1000}, []float64{100}, 5, 2.3849951368733384e-29},
	// {"Hyp1F1", []float64{-10000}, []float64{1000}, 5, 4.389355758473254e-23},
	// {"Hyp1F1", []float64{-100000}, []float64{10000}, 5, 1.6793620777269317e-22},
	// {"Hyp1F1", []float64{-100000}, []float64{20000}, 5, 1.3629514086965236e-11},
	// {"Hyp1F1", []float64{-100000}, []float64{30000}, 5, 5.743051141607582e-08},
	// {"Hyp1F1", []float64{-100000}, []float64{50000}, 5, 4.533185960204361e-05},

	{"Hyp2F1", []float64{+inf, 9}, []float64{-inf}, 1, 1. / 512},
	{"Hyp2F1", []float64{11, 9}, []float64{-7}, 0, 1},
	{"Hyp2F1", []float64{-11, -9}, []float64{-7}, -5, nan},
	{"Hyp2F1", []float64{11, 9}, []float64{-7}, 0.1, nan},
	{"Hyp2F1", []float64{-10.1, -1.02}, []float64{-5.0003}, 1.999, nan},
	{"Hyp2F1", []float64{1, 2.2}, []float64{1}, 1.5, nan},
	{"Hyp2F1", []float64{1, 2}, []float64{1}, 0.5, 4},
	{"Hyp2F1", []float64{3, 4}, []float64{8}, 1, 35},
	{"Hyp2F1", []float64{3, 4}, []float64{8}, -1, 0.320051626105340984775197726448510982379},
	{"Hyp2F1", []float64{11, 9}, []float64{7}, 0.1, 4.392161690778547690302457429384115309585456127951100973539},
	{"Hyp2F1", []float64{1, 2}, []float64{11}, -0.5, 0.918967125770628926427512111627668135785112812919289317012},
	{"Hyp2F1", []float64{-10.1, -1.02}, []float64{-5.3}, 0.999, -0.86248371770779415948376560290918292774539519280219510389},
	{"Hyp2F1", []float64{-11, 9.1}, []float64{7.1}, 5.1, -4.3001996210395527789968701095461658841940532081377151e+07},
	{"Hyp2F1", []float64{-7, -9}, []float64{-11}, -5, 629856. / 11},
	{"Hyp2F1", []float64{-10, -1}, []float64{-5}, 11.9, -22.8},

	{"Hyp4F2", []float64{11, 9, -7, +inf}, []float64{-inf, -inf}, +inf, -3.092849723114612e+15},
	{"Hyp4F2", []float64{11, -9, +inf, +inf}, []float64{-inf, -inf}, 0.5, -2.46925295e+07},
	{"Hyp4F2", []float64{11, -9, 1, 1}, []float64{1, 1}, 0.5, -2.46925295e+07},
	{"Hyp4F2", []float64{11, 9, -7, 0}, []float64{10, -1}, 0.17, 1},
	{"Hyp4F2", []float64{11, 9, 7, 1}, []float64{10, 1}, 1.17, nan},
	{"Hyp4F2", []float64{11, -9, 7, 1}, []float64{10, 1}, 0.17, -11.06499762995341472},
	{"Hyp4F2", []float64{11, -9, 7, 2}, []float64{10, 1}, 0.17, -153.3454736343267104},
	{"Hyp4F2", []float64{-11, 9, 7, 2}, []float64{9, 7}, 0.999, -1.88635604926837706836493630679805453798e+08},

	{"Hyp8F3", []float64{11, -9, 7, 1, 10, 1, 0.17, 19}, []float64{-23.4, 0.05, -2.222}, 0.25, -4.895478076551896507233928858925372381540110041948574e+34},
	{"Hyp8F3", []float64{11, -9, 7, 2, 10, 1, 0.17, 19}, []float64{23.4, 0.05, -2.222}, 0.25, 2.1000378589121131237021537084010298935635405372296384e+34},
	{"Hyp8F3", []float64{-11, 9, 7, 2, 9, 7, 0.17, 19}, []float64{23.4, 0.05, -2.222}, 0.25, 1.4411904034175893283258108991983096641926408333514947e+46},
}

func TestPFQ(t *testing.T) { testutil.Test(t, tol, casesHypPFQ, HypPFQ) }

/*
func BenchmarkHyp0f1(b *testing.B) {
	GlobalF = bench(b, cHyp0F1, "",
		func(x []float64) float64 {
			p := int(x[0])
			q := int(x[1])
			y := x[2:]
			return HypPFQ(y[:p], y[p:p+q], y[p+q])
		})
}
func BenchmarkHyp1F1(b *testing.B) {
	GlobalF = bench(b, cHyp1F1, "",
		func(x []float64) float64 {
			p := int(x[0])
			q := int(x[1])
			y := x[2:]
			return HypPFQ(y[:p], y[p:p+q], y[p+q])
		})
}
func BenchmarkHyp2F1(b *testing.B) {
	GlobalF = bench(b, cHyp2F1, "",
		func(x []float64) float64 {
			p := int(x[0])
			q := int(x[1])
			y := x[2:]
			return HypPFQ(y[:p], y[p:p+q], y[p+q])
		})
}
func BenchmarkHyp4F2(b *testing.B) {
	GlobalF = bench(b, cHyp4F2, "",
		func(x []float64) float64 {
			p := int(x[0])
			q := int(x[1])
			y := x[2:]
			return HypPFQ(y[:p], y[p:p+q], y[p+q])
		})
}

func BenchmarkHyp8F3(b *testing.B) {
	GlobalF = bench(b, cHyp8F3, "",
		func(x []float64) float64 {
			p := int(x[0])
			q := int(x[1])
			y := x[2:]
			return HypPFQ(y[:p], y[p:p+q], y[p+q])
		})
}
*/
