package special_test

import (
	"testing"

	. "github.com/scientificgo/special"
)

var casesLegendreAP = []struct {
	Label    string
	In1, In2 int
	In3, Out float64
}{
	{"", 2, 10, 3.5, 0},
	{"", -10, 10, 3.5, 0},
	{"", 153, 11, 1.1, nan},
	{"", 3, 0, -0.9, -0.4725},
	{"", 10, -3, 0.98, 0.000127413538489068664787307858757370814480523438403071987},
	{"", 22, 10, 3.5, -6.309576310867162196764903515496706859266851097345352e+29},
	{"", 153, 11, 1e-5, 7.1211274720051580808679401818532702913164641927292374e+22},
	{"", -154, 11, 1e-5, 7.1211274720051580808679401818532702913164641927292374e+22},
}

func TestLegendreAP(t *testing.T) {
	for i, c := range casesLegendreAP {
		t.Run(c.Label, func(tt *testing.T) {
			res := LegendreAP(c.In1, c.In2, c.In3)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("[%v]: Got %v, want %v", i, res, c.Out)
			}
		})
	}
}

/*
func BenchmarkLegendreAP(b *testing.B) {
	GlobalF = bench(b, cLegendreAP, "",
		func(x []float64) float64 {
			return LegendreAP(int(x[0]), int(x[1]), x[2])
		})
}
*/
