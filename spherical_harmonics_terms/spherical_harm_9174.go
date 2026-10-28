package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9174 evaluates spherical harmonic radial component order 9174.
func ComputeSphericalHarm9174(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
