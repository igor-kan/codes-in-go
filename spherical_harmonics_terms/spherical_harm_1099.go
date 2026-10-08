package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1099 evaluates spherical harmonic radial component order 1099.
func ComputeSphericalHarm1099(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
