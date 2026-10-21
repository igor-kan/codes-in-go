package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8519 evaluates spherical harmonic radial component order 8519.
func ComputeSphericalHarm8519(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
