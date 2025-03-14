package special_test

import (
	"testing"

	. "github.com/scientificgo/special"
)

var casesHarmonic = []struct {
	Label   string
	In, Out float64
}{
	{"", -inf, nan},
	{"", 0, 0},
	{"", 1, 1},
	{"", 2, 1.5},
	{"", 20, 3.597739657143682},
	{"", 50, 4.499205338329425},
}

func TestHarmonic(t *testing.T) {
	for i, c := range casesHarmonic {
		t.Run(c.Label, func(tt *testing.T) {
			res := Harmonic(c.In)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("[%v]: Got %v, want %v", i, res, c.Out)
			}
		})
	}
}
