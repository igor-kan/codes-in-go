package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1094 evaluates spherical harmonic radial component order 1094.
func ComputeSphericalHarm1094(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
