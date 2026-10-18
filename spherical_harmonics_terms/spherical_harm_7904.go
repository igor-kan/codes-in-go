package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7904 evaluates spherical harmonic radial component order 7904.
func ComputeSphericalHarm7904(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
