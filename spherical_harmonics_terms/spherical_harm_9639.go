package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9639 evaluates spherical harmonic radial component order 9639.
func ComputeSphericalHarm9639(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
