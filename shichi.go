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

func Chi(x float64) float64 {
	return Ei(x) - Shi(x)
}

func Shi(x float64) float64 {
	if x == 0 {
		return 0
	}

	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	if x < 35 {

		z := x * x
		a := 1.0
		s := 1.0
		k := 2.0
		for math.Abs(a/s) > 1e-16 {
			a *= z / k
			k++
			a /= k
			s += a / k
			k++
		}
		return float64(sign) * s * x
	}
	return float64(sign) * Ei(x) / 2
}
