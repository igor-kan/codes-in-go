package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8954 evaluates spherical harmonic radial component order 8954.
func ComputeSphericalHarm8954(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
