package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm89 evaluates spherical harmonic radial component order 89.
func ComputeSphericalHarm89(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
