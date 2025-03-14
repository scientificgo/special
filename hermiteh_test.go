package special_test

import (
	"testing"

	. "github.com/scientificgo/special"
)

var casesHermiteH = []struct {
	Label    string
	In1      int
	In2, Out float64
}{
	{"", 2, nan, nan},
	{"", -2, 2, nan},
	{"", 0, 1.1111, 1},
	{"", 1, 1.1111, 2.2222},
	{"", 4, 5.5, 13201},
	{"", 43, 53.5, 1.56492249523929575819638209988328890215202268133174998288e+87},
}

func TestHermiteH(t *testing.T) {
	for i, c := range casesHermiteH {
		t.Run(c.Label, func(tt *testing.T) {
			res := HermiteH(c.In1, c.In2)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("[%v]: Got %v, want %v", i, res, c.Out)
			}
		})
	}
}

/*
func BenchmarkHermiteH(b *testing.B) {
	GlobalF = bench(b, cHermiteH, "",
		func(x []float64) float64 {
			return HermiteH(int(x[0]), x[1])
		})
}
*/
