/*
   SciGo is a scientific library for the Go language.
   Copyright (C) 2018, Jack Parkinson

   This program is free software: you can redistribute it and/or modify it
   under the terms of the GNU Lesser General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.

   You should have received a copy of the GNU Lesser General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package special

import "math"

// Beta returns the complete beta function, defined by
//
//	Beta(x, y) = Gamma(x) Gamma(y) / Gamma(x+y)
//
// where Gamma is the gamma function.
//
// See http://mathworld.wolfram.com/BetaFunction.html for more information.
func Beta(x, y float64) float64 {
	switch {
	case math.IsNaN(x) || math.IsNaN(y) || math.IsInf(x, -1) || math.IsInf(y, -1):
		return math.NaN()
	case math.IsInf(x, 1):
		if y <= 0 && y == math.Trunc(y) {
			return float64(GammaSign(y)) * x
		}
		return 0
	case math.IsInf(y, 1):
		if x <= 0 && x == math.Trunc(x) {
			return float64(GammaSign(x)) * y
		}
		return 0
	}
	return GammaRatio([]float64{x, y}, []float64{x + y})
}
