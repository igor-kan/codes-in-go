package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm3119 evaluates spherical harmonic radial component order 3119.
func ComputeSphericalHarm3119(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
