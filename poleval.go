// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

// poleval evaluates the polynomial cs[0] + cs[1] * x + ... + cs[n] * x**n
func poleval(x float64, cs ...float64) (p float64) {
	for i := len(cs) - 1; i >= 0; i-- {
		p = cs[i] + p*x
	}
	return
}
