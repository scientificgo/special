package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

var casesGegenbauerC = []struct {
	In1           int
	In2, In3, Out float64
}{
	{2, 1, nan, nan},
	{-2, 1, 2, nan},
	{0, 3.3, 1, 1},
	{1, 1, -3.141, -6.282},
	{1, 2, 3, 12},
	{2, 10, 3.5, 2685},
	{22, 10, 3.5, 5.7734358481154896492250325e+25},
	{40, 15, 4.3255, 6.6613993107922727134428780696365820449481333572147195e+49},
	{40, -15, 4.3255, 0},
	{40, -14.9, 4.3255, -1.280764317783101302989751939635266543812346402681405e+24},
}

func TestGegenbauerC(t *testing.T) {
	for i, c := range casesGegenbauerC {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := GegenbauerC(c.In1, c.In2, c.In3)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
