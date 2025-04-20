package special_test

import (
	"fmt"
	"testing"

	. "github.com/scientificgo/special"
)

func TestSphericalHarmonicY(t *testing.T) {
	cases := []struct {
		In1, In2             int
		In3, In4, Out1, Out2 float64
	}{
		{2, 1, 1, nan, nan, nan},
		{-22, -20, 10, 3.5, nan, nan},
		{0, 0, -7.21, 7.11, 0.282094791773878143474039725780386292922025314664499428422, 0},
		{22, -20, 10, 3.5, 0.00004474740985105335681955012220070375819491035216485506, -0.00005467954127843340910384743380389943632078690198859652},
		{31, 31, -10, 3.5, 5.20105344181401315495645487292339772715804207504464e-10, -4.50058978347382319639863845184135153633008312108735e-09},
		{7, -7, -7, 7, -0.0079410539609713541755773378737370761151396061682892738, -0.025196238025923628008773731309078469019874758134402119},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(tt *testing.T) {
			res1, res2 := SphericalHarmonicY(c.In1, c.In2, c.In3, c.In4)
			ok := equalFloat64(res1, c.Out1) && equalFloat64(res2, c.Out2)
			if !ok {
				tt.Errorf("[%v]: Got (%v, %v), want (%v, %v)", i, res1, res2, c.Out1, c.Out2)
			}
		})
	}
}
