package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9139 evaluates spherical harmonic radial component order 9139.
func ComputeSphericalHarm9139(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
