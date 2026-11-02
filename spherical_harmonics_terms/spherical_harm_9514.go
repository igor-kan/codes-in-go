package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9514 evaluates spherical harmonic radial component order 9514.
func ComputeSphericalHarm9514(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
