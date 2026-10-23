package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8719 evaluates spherical harmonic radial component order 8719.
func ComputeSphericalHarm8719(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
