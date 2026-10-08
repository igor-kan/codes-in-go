package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1119 evaluates spherical harmonic radial component order 1119.
func ComputeSphericalHarm1119(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
