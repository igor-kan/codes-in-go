package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8539 evaluates spherical harmonic radial component order 8539.
func ComputeSphericalHarm8539(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
