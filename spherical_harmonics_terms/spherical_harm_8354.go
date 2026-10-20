package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8354 evaluates spherical harmonic radial component order 8354.
func ComputeSphericalHarm8354(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
