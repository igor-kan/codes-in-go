package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm4004 evaluates spherical harmonic radial component order 4004.
func ComputeSphericalHarm4004(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
