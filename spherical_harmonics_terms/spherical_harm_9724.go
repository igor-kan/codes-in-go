package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9724 evaluates spherical harmonic radial component order 9724.
func ComputeSphericalHarm9724(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
