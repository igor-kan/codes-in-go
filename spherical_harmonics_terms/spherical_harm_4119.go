package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm4119 evaluates spherical harmonic radial component order 4119.
func ComputeSphericalHarm4119(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
