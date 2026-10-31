package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9314 evaluates spherical harmonic radial component order 9314.
func ComputeSphericalHarm9314(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
