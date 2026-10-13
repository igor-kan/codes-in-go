package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7049 evaluates spherical harmonic radial component order 7049.
func ComputeSphericalHarm7049(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
