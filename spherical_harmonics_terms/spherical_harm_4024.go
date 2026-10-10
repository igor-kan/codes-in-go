package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm4024 evaluates spherical harmonic radial component order 4024.
func ComputeSphericalHarm4024(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
