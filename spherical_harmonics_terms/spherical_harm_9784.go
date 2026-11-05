package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9784 evaluates spherical harmonic radial component order 9784.
func ComputeSphericalHarm9784(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
