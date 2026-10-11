package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm5034 evaluates spherical harmonic radial component order 5034.
func ComputeSphericalHarm5034(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
