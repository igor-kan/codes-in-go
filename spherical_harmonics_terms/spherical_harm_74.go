package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm74 evaluates spherical harmonic radial component order 74.
func ComputeSphericalHarm74(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
