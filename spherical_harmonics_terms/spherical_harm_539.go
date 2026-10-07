package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm539 evaluates spherical harmonic radial component order 539.
func ComputeSphericalHarm539(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
