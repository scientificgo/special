// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// gcf evaluates the generalized continued fraction
//
//                         a(1)
//  y = ------------------------------------------
//                            a(2)
//      b(1) + ----------------------------------
//                               a(3)
//             b(2) + --------------------------
//                                  a(4)
//                    b(3) + ------------------
//                                     a(5)
//                           b(4) + ----------
//                                  b(5) +  ⋱
//
func gcf(a, b func(int) float64) float64 {

	// Use the fundamental recurrence formulae
	//
	//  A(i) = b(i)*A1 + a(i)*A2
	//  B(i) = b(i)*B1 + a(i)*A2
	//
	//  y(i) = A(i) / B(i)
	//

	B1, B2 := 1., 0. // B(0) = 1, B(-1) = 0
	A1, A2 := 0., 1. // A(0) = 0, A(-1) = 1

	c := 0. // A1 / B1
	t := inf
	for i := 1; t > macheps; i++ {
		bi := b(i)
		ai := a(i)
		A := bi*A1 + ai*A2
		B := bi*B1 + ai*B2

		r := A / B
		t = math.Abs((c - r) / r)
		c = r

		A1, A2 = A, A1
		B1, B2 = B, B1
	}
	return c
}
