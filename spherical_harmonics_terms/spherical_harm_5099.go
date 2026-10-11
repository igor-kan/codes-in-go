package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm5099 evaluates spherical harmonic radial component order 5099.
func ComputeSphericalHarm5099(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
