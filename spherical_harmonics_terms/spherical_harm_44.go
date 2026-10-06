package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm44 evaluates spherical harmonic radial component order 44.
func ComputeSphericalHarm44(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
