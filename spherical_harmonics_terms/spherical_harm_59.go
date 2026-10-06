package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm59 evaluates spherical harmonic radial component order 59.
func ComputeSphericalHarm59(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
