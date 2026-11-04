package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9644 evaluates spherical harmonic radial component order 9644.
func ComputeSphericalHarm9644(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
