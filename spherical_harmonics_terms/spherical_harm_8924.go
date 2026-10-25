package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8924 evaluates spherical harmonic radial component order 8924.
func ComputeSphericalHarm8924(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
