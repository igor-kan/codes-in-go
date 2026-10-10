package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm4139 evaluates spherical harmonic radial component order 4139.
func ComputeSphericalHarm4139(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
