package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8879 evaluates spherical harmonic radial component order 8879.
func ComputeSphericalHarm8879(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
