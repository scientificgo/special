// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special_test

import "math"

var (
	nan     = math.NaN()
	inf     = math.Inf(1)
	macheps = math.Nextafter(1, 2) - 1
)
