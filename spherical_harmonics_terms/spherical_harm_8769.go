package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8769 evaluates spherical harmonic radial component order 8769.
func ComputeSphericalHarm8769(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
