package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7724 evaluates spherical harmonic radial component order 7724.
func ComputeSphericalHarm7724(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
