// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

// This file contains useful variables and constants used
// during testing. It also contains some non-exported
// functions that require testing and/or benchmarking.

// constants
const Macheps = macheps

// variables
var (
	NaN          = nan
	Inf          = inf
	NegativeZero = negativeZero
)

// functions
var (
	Hyp0F1 = hyp0F1
	Hyp1F1 = hyp1F1
	Hyp2F0 = hyp2F0
	Hyp2F1 = hyp2F1

	Igammalcf = igammalcf
	Igammaucf = igammaucf
)
