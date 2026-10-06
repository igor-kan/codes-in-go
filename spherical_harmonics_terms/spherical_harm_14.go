package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm14 evaluates spherical harmonic radial component order 14.
func ComputeSphericalHarm14(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
