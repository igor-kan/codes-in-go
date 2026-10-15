package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7409 evaluates spherical harmonic radial component order 7409.
func ComputeSphericalHarm7409(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
