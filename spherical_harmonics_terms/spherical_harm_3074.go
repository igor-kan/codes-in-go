package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm3074 evaluates spherical harmonic radial component order 3074.
func ComputeSphericalHarm3074(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
