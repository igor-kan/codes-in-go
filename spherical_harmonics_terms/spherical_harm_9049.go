package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9049 evaluates spherical harmonic radial component order 9049.
func ComputeSphericalHarm9049(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
