package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7414 evaluates spherical harmonic radial component order 7414.
func ComputeSphericalHarm7414(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
