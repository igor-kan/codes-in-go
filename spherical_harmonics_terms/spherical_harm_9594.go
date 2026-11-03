package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9594 evaluates spherical harmonic radial component order 9594.
func ComputeSphericalHarm9594(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
