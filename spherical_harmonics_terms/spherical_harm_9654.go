package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9654 evaluates spherical harmonic radial component order 9654.
func ComputeSphericalHarm9654(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
