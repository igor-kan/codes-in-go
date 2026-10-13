package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7034 evaluates spherical harmonic radial component order 7034.
func ComputeSphericalHarm7034(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
