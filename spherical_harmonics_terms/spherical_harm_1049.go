package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1049 evaluates spherical harmonic radial component order 1049.
func ComputeSphericalHarm1049(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
