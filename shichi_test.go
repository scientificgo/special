package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

func TestShi(t *testing.T) {
	cases := []struct {
		In, Out float64
	}{
		{0, 0},
		{2, 2.5015674333549756},
		{7.9, 202.35036897887048},
		{-7.9, -202.35036897887048},
		{9, 518.9391515822219},
		{10, 1246.1144901994232},
		{20, 1.2807826332028294e+07},
		{80, 3.5073000024524e+32},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := Shi(c.In)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}

func TestChi(t *testing.T) {
	cases := []struct {
		In, Out float64
	}{
		{20, 1.28078263320282943610629339487996274627064136343962909e+07},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := Chi(c.In)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
