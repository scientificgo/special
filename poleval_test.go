// Copyright (c) 2020, Jack Parkinson. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package special

import "testing"

func TestPoleval(t *testing.T) {
	type args struct {
		x  float64
		cs []float64
	}
	tests := []struct {
		name  string
		args  args
		wantP float64
	}{
		{"x=0", args{0, []float64{1}}, 1},
		{"x=0", args{0, []float64{2.2, 1, 1}}, 2.2},
		{"x=1", args{1, []float64{2.2, 1, 1}}, 4.2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotP := Poleval(tt.args.x, tt.args.cs...); gotP != tt.wantP {
				t.Errorf("Poleval() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}
