package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8349 evaluates spherical harmonic radial component order 8349.
func ComputeSphericalHarm8349(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
