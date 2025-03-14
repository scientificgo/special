package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

var casesFibonacci = []struct {
	In, Out float64
}{
	{+inf, +inf},
	{-inf, nan},
	{nan, nan},
	{0, 0},
	{1, 1},
	{2, 1},
	{-1, 1},
	{-2, -1},
	{2.5, 1.4893065462657091},
	{-123.321, 1.4123645428734893e+25},
	{-100, -3.542248481792619e+20},
	{1.618033988749895, 0.8998442262232443},
	{87, 6.798916376386122e+17},
}

func TestFibonacci(t *testing.T) {
	for i, c := range casesFibonacci {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res := Fibonacci(c.In)
			ok := equalFloat64(res, c.Out)
			if !ok {
				tt.Errorf("Got %v, want %v", res, c.Out)
			}
		})
	}
}
