package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8869 evaluates spherical harmonic radial component order 8869.
func ComputeSphericalHarm8869(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
