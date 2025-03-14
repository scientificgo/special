package special_test

import (
	"testing"

	. "github.com/scientificgo/special"
)

var casesDigamma = []struct {
	Label   string
	In, Out float64
}{
	{"", nan, nan},
	{"", -inf, nan},
	{"", +inf, +inf},
	{"", -2., nan},
	{"", 1e+08, 18.420680738952367},
	{"", 1e-08, -1.0000000057721564e+08},
	{"", -1e-08, 9.999999942278431e+07},
	{"", 10., 2.251752589066721},
	{"", -10.2, 6.6946384504174965},
	{"", 4.9, 1.483737793254897},
	{"", -4.9, -7.981008564556067},
}

func TestDigamma(t *testing.T) {
	for i, c := range casesDigamma {
		t.Run(c.Label, func(tt *testing.T) {
			res := Digamma(c.In)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("[%v]: Got %v, want %v", i, res, c.Out)
			}
		})
	}
}

/*
func BenchmarkDigamma(b *testing.B) {
	GlobalF = bench(b, cDigamma, "",
		func(x []float64) float64 {
			return Digamma(x[0])
		})
}
*/
