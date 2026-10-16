package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7624 evaluates spherical harmonic radial component order 7624.
func ComputeSphericalHarm7624(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
