package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9714 evaluates spherical harmonic radial component order 9714.
func ComputeSphericalHarm9714(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
