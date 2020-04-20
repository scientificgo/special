// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

// stirling evaluates the first seven terms of the series
// in Stirling's approximation for the Gamma function
// multiplied by sqrt(2*π).
//
//  Gamma(x) ~ (x/e)**x / sqrt(x) * stirling(x)
//
func stirling(x float64) float64 {
	_stirling := []float64{
		1.,
		8.33333333333482257126e-02,
		3.47222221605458667310e-03,
		-2.68132617805781232825e-03,
		-2.29549961613378126380e-04,
		7.87311395793093628397e-04,
		// 6.97281375836585777429e-05, // does not seem to improve accuracy
	}
	return Sqrt2Pi * poleval(1/x, _stirling...)
}
