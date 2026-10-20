package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8314 evaluates spherical harmonic radial component order 8314.
func ComputeSphericalHarm8314(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
