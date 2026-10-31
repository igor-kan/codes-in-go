package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9379 evaluates spherical harmonic radial component order 9379.
func ComputeSphericalHarm9379(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
