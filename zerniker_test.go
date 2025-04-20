package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

func TestZernikeR(t *testing.T) {
	cases := []struct {
		In1, In2 int
		In3, Out float64
	}{
		{-1, 20, 1.1, nan},
		{987, 988, 98765432.1234567, 0},
		{7, 5, 3.3, 27484.7865039},
		{43, 41, 53.5, 8.9646400010525796235939120335742702428482112823125695e+75},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := ZernikeR(c.In1, c.In2, c.In3)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
