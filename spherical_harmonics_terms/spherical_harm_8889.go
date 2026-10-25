package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8889 evaluates spherical harmonic radial component order 8889.
func ComputeSphericalHarm8889(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
