package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm549 evaluates spherical harmonic radial component order 549.
func ComputeSphericalHarm549(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
