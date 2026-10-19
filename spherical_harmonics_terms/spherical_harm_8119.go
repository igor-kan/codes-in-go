package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8119 evaluates spherical harmonic radial component order 8119.
func ComputeSphericalHarm8119(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
