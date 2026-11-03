package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9574 evaluates spherical harmonic radial component order 9574.
func ComputeSphericalHarm9574(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
