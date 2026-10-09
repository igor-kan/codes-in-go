package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm3024 evaluates spherical harmonic radial component order 3024.
func ComputeSphericalHarm3024(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
