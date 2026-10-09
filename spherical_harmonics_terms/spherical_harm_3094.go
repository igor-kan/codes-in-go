package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm3094 evaluates spherical harmonic radial component order 3094.
func ComputeSphericalHarm3094(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
