// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

// Package special contains special mathematical functions
// and constants to supplement the standard math package.
package special // import "scientificgo.org/special"

import "math"

// Mathematical constants.
const (
	Euler = 0.577215664901532860606512090082402431042159335939923598805767234884867726777664670936947063291746749 // https://oeis.org/A001620

	LnPi  = 1.144729885849400174143427351353058711647294812915311571513623071472137769884826079783623270275489707 // https://oeis.org/A053510
	LnPhi = 0.481211825059603447497758913424368423135184334385660519661018168840163867608221774412009429122723474 // https://oeis.org/A002390

	Sqrt2Pi = 2.506628274631000502415765284811045253006986740609938316629923576342293654607841974946595838378057266 // https://oeis.org/A019727
)

const macheps = 2. / (1 << 53) // machine epsilon, or ε. Numerically equivalent to math.Nextafter(1, 2) - 1

var (
	nan = math.NaN()
	inf = math.Inf(1)

	negativeZero = math.Copysign(0, -1)
)
