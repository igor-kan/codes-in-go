package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm5029 evaluates spherical harmonic radial component order 5029.
func ComputeSphericalHarm5029(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
