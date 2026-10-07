package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm504 evaluates spherical harmonic radial component order 504.
func ComputeSphericalHarm504(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
