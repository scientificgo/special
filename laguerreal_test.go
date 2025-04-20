package special_test

import (
	"fmt"
	"math"
	"testing"

	. "github.com/scientificgo/special"
)

func TestLaguerreAL(t *testing.T) {
	cases := []struct {
		In1           int
		In2, In3, Out float64
	}{
		{2, 1, nan, nan},
		{-2, 1, 2, nan},
		{0, 22, 9.98, 1},
		{0, 22, +inf, 1},
		{0, 22, -inf, 1},
		{1, +inf, +inf, nan},
		{1, math.Sqrt2, math.Sqrt2, 1},
		{1, 2, -math.Pi, 1 + 2 + math.Pi},
		{1, 1, -inf, +inf},
		{41, -47, 4.3255, nan},
		{2, -1, 3.5, 2.625},
		{2, -1.5, 3.5, 4.25},
		{20, -math.Pi, math.Phi, -0.0115694547113679203779196470133707831118560425919161933516912962215768109},
		{20, -math.Pi, -math.Phi, 47.1942974528621228434374048263934823058447034757284757020894529377955060066},
		{22, 10, 3.5, 25252.39978746937700539858666621210279857995048910207458046},
		{41, 17, -4.3255, 4.2086917252857569297940591579340234265295868262479278e+17},
		{41, -17, 4.3255, 4.9579126998422110619614409645401386382056238641622942e-08},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := LaguerreAL(c.In1, c.In2, c.In3)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
