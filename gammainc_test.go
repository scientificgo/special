package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

func TestGammaIncU(t *testing.T) {
	var casesGammaIncU = []struct {
		In1, In2, Out float64
	}{
		{nan, 2, nan},
		{20, -2.432, nan},
		{10, 0, 362880},
		{10, +inf, 0},
		{0, 10, 4.156968929685325e-06},
		{-1, 10, 3.830240465631609e-07},
		{-10, 10, 2.2146903192202743e-16},
		{-10, 27, 2.420076067270557e-28},
		{10, 10, 166173.53478754574},
		{10, 1, 362879.95956592244},
		{-10.2, 1.99, 9.893689107832149e-06},
	}
	for i, c := range casesGammaIncU {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := GammaIncU(c.In1, c.In2)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}

func TestGammaIncL(t *testing.T) {
	var casesGammaIncL = []struct {
		In1, In2, Out float64
	}{
		{nan, 2, nan},
		{-10, nan, nan},
		{20, -2.432, nan},
		{10, 0, 0},
		{10, +inf, 362880},
		{-10, 10, +inf},
		{10, 100, 362880},
		{10, 1000, 362880},
		{100, 1000, 9.332621544394415e+155},
		{10, 0.001, 9.990913256294004e-32},
		{-11.2, 1.99, -4.522214610043099e-06},
	}
	for i, c := range casesGammaIncL {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := GammaIncL(c.In1, c.In2)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}

func TestGammaIncIdentity(t *testing.T) {
	cases := []struct {
		In1, In2, Out float64
	}{
		{nan, 2, nan},
		{-10, nan, nan},
		{20, -2.432, nan},
		{+inf, 12.2, nan},
		{10, +inf, 362880},
		{10, 43, 362880},
		{50.5, 94.3, 4.290462912351957e+63},
		{150.5, 94.3, 4.661072627097374e+261},
		{150.5, 194.3, 4.661072627097374e+261},
		{150.5, 1.943e+07, 4.661072627097374e+261},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := GammaIncU(c.In1, c.In2) + GammaIncL(c.In1, c.In2)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
