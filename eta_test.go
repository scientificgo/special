package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

var casesEta = []struct {
	In, Out float64
}{
	{nan, nan},
	{-1, 0.25},
	{0, 0.5},
	{1, 0.6931471805599453},
	{+inf, 1},
	{-2, 0},
	{-1000, 0},
	{-1.1, 0.2234614116270079},
	{1.1, 0.7088088499305867},
}

func TestEta(t *testing.T) {
	for i, c := range casesEta {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := Eta(c.In)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
