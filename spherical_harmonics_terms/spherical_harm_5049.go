package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm5049 evaluates spherical harmonic radial component order 5049.
func ComputeSphericalHarm5049(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
