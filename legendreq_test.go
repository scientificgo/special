package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

var casesLegendreQ = []struct {
	In1      int
	In2, Out float64
}{
	{2, nan, nan},
	{-2, 2, nan},
	{1, 1.1, nan},
	{0, 0.9, 1.472219489583220230004513715943926768618689630649564409268},
	{1, 0.999, 2.796400966082949831744191300541195457801412019535269391030},
	{11, 0.999, 0.665248555792627905833229643332143091673876114974304501168},
	{101, -0.10101, 0.082745695703743357501272084016026991789721245950554478861},
}

func TestLegendreQ(t *testing.T) {
	for i, c := range casesLegendreQ {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := LegendreQ(c.In1, c.In2)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
