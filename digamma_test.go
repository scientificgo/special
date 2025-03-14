package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

var casesDigamma = []struct {
	In, Out float64
}{
	{nan, nan},
	{-inf, nan},
	{+inf, +inf},
	{-2., nan},
	{1e+08, 18.420680738952367},
	{1e-08, -1.0000000057721564e+08},
	{-1e-08, 9.999999942278431e+07},
	{10., 2.251752589066721},
	{-10.2, 6.6946384504174965},
	{4.9, 1.483737793254897},
	{-4.9, -7.981008564556067},
}

func TestDigamma(t *testing.T) {
	for i, c := range casesDigamma {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := Digamma(c.In)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
