package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8514 evaluates spherical harmonic radial component order 8514.
func ComputeSphericalHarm8514(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
