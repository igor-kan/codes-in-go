package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8834 evaluates spherical harmonic radial component order 8834.
func ComputeSphericalHarm8834(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
