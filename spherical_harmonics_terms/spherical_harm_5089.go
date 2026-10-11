package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm5089 evaluates spherical harmonic radial component order 5089.
func ComputeSphericalHarm5089(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
