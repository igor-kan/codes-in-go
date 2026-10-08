package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1089 evaluates spherical harmonic radial component order 1089.
func ComputeSphericalHarm1089(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
