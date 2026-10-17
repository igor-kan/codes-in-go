package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7744 evaluates spherical harmonic radial component order 7744.
func ComputeSphericalHarm7744(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
