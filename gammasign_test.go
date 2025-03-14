package special_test

import (
	"testing"

	. "github.com/scientificgo/special"
)

var casesGammaSign = []struct {
	Label string
	In    float64
	Out   int
}{
	{"", +inf, 1},
	{"", -inf, 1},
	{"", nan, 1},
	{"", 0.0, 1},
	{"", 1.0, 1},
	{"", -1.0, -1},
	{"", -2.0, 1},
}

func TestGammaSign(t *testing.T) {
	for i, c := range casesGammaSign {
		t.Run(c.Label, func(tt *testing.T) {
			res := GammaSign(c.In)
			ok := res == c.Out
			if !ok {
				tt.Errorf("[%v]: Got %v, want %v", i, res, c.Out)
			}
		})
	}
}

/*
func BenchmarkGammaSign(b *testing.B) {
	GlobalF = bench(b, cGammaSign, "",
		func(x []float64) float64 {
			return float64(GammaSign(x[0]))
		})
}
*/
