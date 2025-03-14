package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

var casesTrigamma = []struct {
	In, Out float64
}{
	{nan, nan},
	{-inf, nan},
	{+inf, 0.},
	{-2., nan},
	{1e+08, 1.000000005e-08},
	{1e-08, 1.0000000000000002e+16},
	{-1e-08, 1.0000000000000002e+16},
	{10., 0.10516633568168575},
	{-10.2, 28.473461217287777},
	{4.9, 0.22631146419168002},
	{-4.9, 103.17117695886462},
}

func TestTrigamma(t *testing.T) {
	for i, c := range casesTrigamma {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := Trigamma(c.In)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
