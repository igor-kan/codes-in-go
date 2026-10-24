package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8844 evaluates spherical harmonic radial component order 8844.
func ComputeSphericalHarm8844(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
