package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7204 evaluates spherical harmonic radial component order 7204.
func ComputeSphericalHarm7204(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
