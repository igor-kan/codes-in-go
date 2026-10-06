package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm109 evaluates spherical harmonic radial component order 109.
func ComputeSphericalHarm109(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
