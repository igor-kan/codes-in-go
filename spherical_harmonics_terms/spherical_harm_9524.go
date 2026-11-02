package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9524 evaluates spherical harmonic radial component order 9524.
func ComputeSphericalHarm9524(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
