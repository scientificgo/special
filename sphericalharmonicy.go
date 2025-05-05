package special

import "math"

// SphericalHarmonicY returns the angular portion of the solutions to Laplace's equation
// in spherical coordinates, where theta is in [0, π], phi is in [0, 2π], l ≥ 0, |m| ≤ l.
//
// See http://mathworld.wolfram.com/SphericalHarmonic.html for more information.
func SphericalHarmonicY(l int, m int, theta float64, phi float64) (float64, float64) {
	if math.IsNaN((theta)) || math.IsNaN(phi) {
		return math.NaN(), math.NaN()
	}

	if l < 0 || m > l || -m > l {
		return math.NaN(), math.NaN()
	}

	// TODO: is this special case necessary?
	if m == -l {
		st := math.Sin(theta)
		im, re := math.Sincos(phi * float64(-m))
		s := math.Sqrt(math.Gamma(2*float64(l+1))) / math.Gamma(float64(l+1)) * math.Pow(st/2, float64(l)) / (2 * math.SqrtPi)
		return re * s, -im * s
	}

	ct := math.Cos(theta)
	p := math.Sqrt(float64(2*l+1) * math.Gamma(float64(l-m+1)) / math.Gamma(float64(l+m+1)))
	s := p * LegendreAP(l, m, ct) / (2 * math.SqrtPi)
	im, re := math.Sincos(phi * float64(m))

	return re * s, im * s
}
