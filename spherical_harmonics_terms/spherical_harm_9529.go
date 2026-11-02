package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9529 evaluates spherical harmonic radial component order 9529.
func ComputeSphericalHarm9529(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
