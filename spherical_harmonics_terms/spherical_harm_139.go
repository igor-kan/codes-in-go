package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm139 evaluates spherical harmonic radial component order 139.
func ComputeSphericalHarm139(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
