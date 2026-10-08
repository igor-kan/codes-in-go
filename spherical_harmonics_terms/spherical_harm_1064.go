package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1064 evaluates spherical harmonic radial component order 1064.
func ComputeSphericalHarm1064(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
