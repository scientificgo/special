// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import (
	"fmt"
	"math"

	. "scientificgo.org/special"
)

// log1p returns Log(1+x/y) using the generalized continued
// fraction representation, given by
//                                1
//  ln(1+x/y)/x = ----------------------------------
//                                 1x
//	              1y + ----------------------------
//                                   1x
//                    2 + ------------------------
//                                     2x
//                        3y + ------------------
//                                       2x
//                             2 + -------------
//                                        3x
//                                 5y + -------
//                                      2 + ...
//
func log1p(x, y float64) float64 {
	// a(i) = floor(i/2) * x
	a := func(i int) float64 {
		return float64(i/2) * x // integer division automatically floors
	}
	// b(i) = y*i for odd i, else 2
	b := func(i int) float64 {
		if i%2 == 1 {
			return float64(i) * y
		}
		return 2
	}
	return x * GCF(a, b)
}

func ExampleGCF_log1p() {
	x := 2.25
	y := 0.01125
	fmt.Println(log1p(x, y) == math.Log1p(x/y))
	// Output: true
}
