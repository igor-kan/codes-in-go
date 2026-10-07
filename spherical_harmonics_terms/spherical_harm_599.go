package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm599 evaluates spherical harmonic radial component order 599.
func ComputeSphericalHarm599(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
