package special_test

import (
	"testing"

	. "github.com/scientificgo/special"
)

var casesTrigamma = []struct {
	Label   string
	In, Out float64
}{
	{"", nan, nan},
	{"", -inf, nan},
	{"", +inf, 0.},
	{"", -2., nan},
	{"", 1e+08, 1.000000005e-08},
	{"", 1e-08, 1.0000000000000002e+16},
	{"", -1e-08, 1.0000000000000002e+16},
	{"", 10., 0.10516633568168575},
	{"", -10.2, 28.473461217287777},
	{"", 4.9, 0.22631146419168002},
	{"", -4.9, 103.17117695886462},
}

func TestTrigamma(t *testing.T) {
	for i, c := range casesTrigamma {
		t.Run(c.Label, func(tt *testing.T) {
			res := Trigamma(c.In)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("[%v]: Got %v, want %v", i, res, c.Out)
			}
		})
	}
}

/*
func BenchmarkTrigamma(b *testing.B) {
	GlobalF = bench(b, cTrigamma, "",
		func(x []float64) float64 {
			return Trigamma(x[0])
		})
}
*/
