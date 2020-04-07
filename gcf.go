// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "math"

// GCF evaluates the generalized continued fraction
// defined by the functions a(i≥2) and b(i≥1), which
// return the i-th partial numerator and partial denominator
// respectively. It assumes, wlog, that b(0)=0 and a(1)=1; other
// values can be trivially recovered as y' = b(0) + a(1)*y.
//
//                          1
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
func GCF(a, b func(int) float64) float64 {

	// Use the fundamental recurrence formulae
	//
	//  A(i) = b(i)*A(i-1) + a(i)*A(i-2)
	//  B(i) = b(i)*B(i-1) + a(i)*A(i-2)
	//
	//  y(i) = A(i) / B(i)
	//
	// until |1 - y(n-1)/y(n)| < ε
	//

	const big = 1e200

	A1, A2 := 1., 0.
	B1, B2 := b(1), 1.

	y := A1 / B1 // y(1) = A(1) / B(1)
	t := inf
	for i := 2; t > macheps; i++ { // iterate to get A(2), B(2) and y(2) etc.
		ai := a(i)
		bi := b(i)
		A := bi*A1 + ai*A2
		B := bi*B1 + ai*B2

		r := A / B
		t = math.Abs((y - r) / y)
		y = r

		A1, A2 = A, A1
		B1, B2 = B, B1

		if A1 > big { // rescale to avoid overflows
			A1 /= big
			A2 /= big
			B1 /= big
			B2 /= big
		}
	}
	return y
}
