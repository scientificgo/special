package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

func TestBeta(t *testing.T) {
	cases := []struct {
		In1, In2, Out float64
	}{
		{nan, 1, nan},
		{-inf, -inf, nan},
		{+inf, +inf, 0},
		{+inf, 190.7, 0},
		{0.986, -inf, nan},
		{-3, -7, +inf},
		{10, -10, 0.1},
		{-15, 15, -0.06666666666666667},
		{1.5, -2.5, 0},
		{3.54, -1, -inf},
		{-1, 3.54, -inf},
		{0, 12, +inf},
		{+inf, -7, -inf},
		{-8, +inf, +inf},
		{8, +inf, 0},
		{300, 200, 1.6485491608664747e-147},
		{0.5, -0.75, 1.74803836952808},
		{100000, 1e-05, 99987.91060292121},
		{-1.123, -1.132, -33.830986471614295},
		{-112.3, -113.2, -2.668986182849379e+67},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := Beta(c.In1, c.In2)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
