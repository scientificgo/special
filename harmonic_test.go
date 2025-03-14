package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

var casesHarmonic = []struct {
	In, Out float64
}{
	{-inf, nan},
	{0, 0},
	{1, 1},
	{2, 1.5},
	{20, 3.597739657143682},
	{50, 4.499205338329425},
}

func TestHarmonic(t *testing.T) {
	for i, c := range casesHarmonic {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := Harmonic(c.In)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
