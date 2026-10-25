package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8969 evaluates spherical harmonic radial component order 8969.
func ComputeSphericalHarm8969(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
