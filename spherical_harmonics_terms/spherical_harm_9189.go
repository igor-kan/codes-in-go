package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9189 evaluates spherical harmonic radial component order 9189.
func ComputeSphericalHarm9189(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
