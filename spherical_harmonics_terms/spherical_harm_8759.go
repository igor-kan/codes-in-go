package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8759 evaluates spherical harmonic radial component order 8759.
func ComputeSphericalHarm8759(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
