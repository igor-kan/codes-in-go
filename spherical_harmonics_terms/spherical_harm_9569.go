package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9569 evaluates spherical harmonic radial component order 9569.
func ComputeSphericalHarm9569(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
