package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9024 evaluates spherical harmonic radial component order 9024.
func ComputeSphericalHarm9024(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
