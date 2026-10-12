package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm6034 evaluates spherical harmonic radial component order 6034.
func ComputeSphericalHarm6034(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
