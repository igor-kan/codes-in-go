package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm4029 evaluates spherical harmonic radial component order 4029.
func ComputeSphericalHarm4029(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
