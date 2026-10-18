package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7929 evaluates spherical harmonic radial component order 7929.
func ComputeSphericalHarm7929(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
