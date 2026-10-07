package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm569 evaluates spherical harmonic radial component order 569.
func ComputeSphericalHarm569(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
