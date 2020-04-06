// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// Li returns the logarithmic integral of x.
//
// Special cases are:
//  Li(0) = 0
//  Li(1) = -Inf
//  Li(+Inf) = +Inf
//  Li(x < 0) = NaN
//
func Li(x float64) float64 { return Ei(math.Log(x)) }
