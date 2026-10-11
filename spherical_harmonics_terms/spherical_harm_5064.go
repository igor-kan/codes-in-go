package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm5064 evaluates spherical harmonic radial component order 5064.
func ComputeSphericalHarm5064(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
