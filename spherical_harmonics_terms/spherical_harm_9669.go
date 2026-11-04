package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9669 evaluates spherical harmonic radial component order 9669.
func ComputeSphericalHarm9669(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
