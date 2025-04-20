package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

func TestLi(t *testing.T) {
	cases := []struct {
		In, Out float64
	}{
		{-1, nan},
		{1e-300, -1.4455586289196509e-303},
		{1e-100, -4.324245377202095e-103},
		{1e-50, -8.611722553713826e-53},
		{1e-25, -1.707997398334023e-27},
		{1e-15, -2.815944392958954e-17},
		{1e-14, -3.0113475954567504e-16},
		{1e-13, -3.2359176090698706e-15},
		{1e-12, -3.4967184067697225e-14},
		{1e-10, -4.168887750019648e-12},
		{1e-08, -5.161659103222967e-10},
		{1e-07, -5.859859508790472e-09},
		{1e-05, -8.036767291447334e-07},
		{2, 1.045163780117493},
		{1e+250, 1.740206254656917e+247},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := Li(c.In)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
