package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9169 evaluates spherical harmonic radial component order 9169.
func ComputeSphericalHarm9169(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
