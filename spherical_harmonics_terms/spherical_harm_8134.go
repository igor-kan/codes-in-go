package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8134 evaluates spherical harmonic radial component order 8134.
func ComputeSphericalHarm8134(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
