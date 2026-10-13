package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7004 evaluates spherical harmonic radial component order 7004.
func ComputeSphericalHarm7004(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
