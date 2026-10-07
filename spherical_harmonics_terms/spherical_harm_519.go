package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm519 evaluates spherical harmonic radial component order 519.
func ComputeSphericalHarm519(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
