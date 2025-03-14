package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

var casesGammaSign = []struct {
	In  float64
	Out int
}{
	{+inf, 1},
	{-inf, 1},
	{nan, 1},
	{0.0, 1},
	{1.0, 1},
	{-1.0, -1},
	{-2.0, 1},
}

func TestGammaSign(t *testing.T) {
	for i, c := range casesGammaSign {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := GammaSign(c.In)
			ok := res == c.Out
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
