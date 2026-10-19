package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8124 evaluates spherical harmonic radial component order 8124.
func ComputeSphericalHarm8124(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
