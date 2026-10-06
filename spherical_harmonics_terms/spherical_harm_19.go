package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm19 evaluates spherical harmonic radial component order 19.
func ComputeSphericalHarm19(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
