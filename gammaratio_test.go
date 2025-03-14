package special_test

import (
	"testing"

	. "github.com/scientificgo/special"
)

var casesGammaRatio = []struct {
	Label    string
	In1, In2 []float64
	Out      float64
}{
	{"", []float64{18.3}, []float64{17.3}, 17.3},
	{"", []float64{18.3}, []float64{19.3}, 0.054644808743169398907103825136612021857923497267759562841},
	{"", []float64{17.3}, []float64{18.3}, 0.057803468208092485549132947976878612716763005780346820809},
	{"", []float64{3056.675}, []float64{2994.86}, 1.455279029640759935931407996328976678151e+215},
	{"", []float64{52}, []float64{34}, 1.78632041505308592071270400000e+29},
	{"", []float64{5200}, []float64{5170}, 2.7617744500677061199284498641186414738660945163980473e+111},
	{"", []float64{-3056.675}, []float64{-2994.86}, 3.36201570269494877318076138295338657248e-216},
	{"", []float64{30.56675}, []float64{-29.9486}, 6.8925568889573652010911108843213060071431578071870857e+62},
	{"", []float64{-30.56675}, []float64{29.9486}, -2.336633827644530886362525262834410927501228572491793e-64},
	{"", []float64{-30.56675}, []float64{2.99486}, -8.722212449104196743424739522411319053538568399884817e-34},
	{"", []float64{30.56675}, []float64{-2.99486}, -1.854447343857626562702508228156158175442536497190872e+30},
	{"", []float64{-30.56675}, []float64{-2.99486}, 5.3197612610626883363834400926838090070275272412206817e-35},
	{"", []float64{-3.056675}, []float64{-2.99486}, -0.08432044283837548833933209111172014690828397293968886059},
	{"", []float64{3.056675}, []float64{-2.99486}, -0.06461147512228474627478404097611852893360651779019779098},
	{"", []float64{-3.056675}, []float64{2.99486}, 1.382507184339208614977034839352802742244932586022841773899},
	{"", []float64{-3.056675e-05}, []float64{-2.99486e-05}, 0.979777394939253166323172227477217052309400947358568028795},
	{"", []float64{-3.056675e-05}, []float64{2.99486e-05}, -0.97981126995868406921541291476627784618195757728516169958},
}

func TestGammaRatio(t *testing.T) {
	for i, c := range casesGammaRatio {
		t.Run(c.Label, func(tt *testing.T) {
			res := GammaRatio(c.In1, c.In2)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("[%v]: Got %v, want %v", i, res, c.Out)
			}
		})
	}
}
