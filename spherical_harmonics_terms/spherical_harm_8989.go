package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8989 evaluates spherical harmonic radial component order 8989.
func ComputeSphericalHarm8989(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
