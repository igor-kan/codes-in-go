package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9094 evaluates spherical harmonic radial component order 9094.
func ComputeSphericalHarm9094(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
