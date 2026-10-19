package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8164 evaluates spherical harmonic radial component order 8164.
func ComputeSphericalHarm8164(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
