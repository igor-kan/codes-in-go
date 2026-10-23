package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8774 evaluates spherical harmonic radial component order 8774.
func ComputeSphericalHarm8774(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
