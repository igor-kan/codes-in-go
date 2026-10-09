package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm3089 evaluates spherical harmonic radial component order 3089.
func ComputeSphericalHarm3089(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
