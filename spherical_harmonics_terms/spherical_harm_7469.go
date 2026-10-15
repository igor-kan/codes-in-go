package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7469 evaluates spherical harmonic radial component order 7469.
func ComputeSphericalHarm7469(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
