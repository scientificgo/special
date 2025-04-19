package special

import (
	"fmt"
	"math"
	"testing"
)

func TestRemoveCommonElements(t *testing.T) {
	cases := []struct {
		In1, In2   []float64
		Out1, Out2 []float64
	}{
		{
			[]float64{}, []float64{},
			[]float64{}, []float64{},
		},
		{
			[]float64{1}, []float64{2},
			[]float64{1}, []float64{2},
		},
		{
			[]float64{1, 2, 3}, []float64{1, 2, 4, 5},
			[]float64{3}, []float64{4, 5},
		},
		{
			[]float64{1, 2, 3, math.NaN(), math.Inf(1)}, []float64{math.NaN(), math.Inf(1), 4, 5, 6},
			[]float64{1, 2, 3, math.NaN()}, []float64{math.NaN(), 4, 5, 6},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res1, res2, res3, res4 := removeCommonElements(c.In1, c.In2)
			ok := res3 == len(c.Out1) && res4 == len(c.Out2)
			if ok {
				for j, r := range res1 {
					eq := r == c.Out1[j] || math.IsNaN(r) == math.IsNaN(c.Out1[j])
					if !eq {
						ok = false
						break
					}
				}
				for j, r := range res2 {
					eq := r == c.Out2[j] || math.IsNaN(r) == math.IsNaN(c.Out2[j])
					if !eq {
						ok = false
						break
					}
				}
			}
			if !ok {
				tt.Errorf("Got (%v, %v), want (%v, %v)", res1, res2, c.Out1, c.Out2)
			}
		})
	}
}

func TestPoly(t *testing.T) {
	cases := []struct {
		In1 float64
		In2 []float64
		Out float64
	}{
		{0, []float64{1, 2, 3, 4, 5}, 1},
		{1, []float64{1, 2, 3, 4, 5}, 15},
		{2, []float64{1}, 1},
		{2, []float64{1, 2}, 1 + 2*2},
		{2, []float64{1, 2, 3}, 1 + 2*2 + 3*4},
		{2, []float64{1, 2, 3, 4}, 1 + 2*2 + 3*4 + 4*8},
		{2, []float64{1, 2, 3, 4, 5}, 1 + 2*2 + 3*4 + 4*8 + 5*16},
		{math.Pi, []float64{math.Sqrt2, math.SqrtE, math.SqrtE}, math.Sqrt2 + math.SqrtE*math.Pi + math.SqrtE*math.Pi*math.Pi},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := poly(c.In1, c.In2...)
			ok := res == c.Out
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
