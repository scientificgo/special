package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

var casesLegendreP = []struct {
	In1      int
	In2, Out float64
}{
	{2, nan, nan},
	{0, 1.21, 1},
	{1, -3.141, -3.141},
	{153, 1e-5, 0.000098854224351825737394494825999239936670301828815467202},
	{-2, -3.141, -3.141},
}

func TestLegendreP(t *testing.T) {
	for i, c := range casesLegendreP {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := LegendreP(c.In1, c.In2)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
