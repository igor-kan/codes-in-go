package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9754 evaluates spherical harmonic radial component order 9754.
func ComputeSphericalHarm9754(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
