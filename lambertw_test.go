package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

var casesLambertW = []struct {
	In1      int
	In2, Out float64
}{
	{2, 2.2, nan},
	{0, 0, 0},
	{0, -3, nan},
	{0, +inf, +inf},
	{-1, 0, -inf},
	{-1, -3, nan},
	{-1, 3, nan},
	{0, -0.33, -0.6032666497551331},
	{0, 0.1, 0.09127652716086226},
	{0, 4.5, 1.2672378143074348},
	{0, 9.9, 1.7391425517333516},
	{0, 10.89, 1.8000374607381258},
	{0, 100.12, 3.386555992882349},
	{-1, -0.36787944117144233, -1},
	{-1, -0.33, -1.541268224332639},
	{-1, -0.1, -3.577152063957297},
}

func TestLambertW(t *testing.T) {
	for i, c := range casesLambertW {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := LambertW(c.In1, c.In2)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
