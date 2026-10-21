package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8534 evaluates spherical harmonic radial component order 8534.
func ComputeSphericalHarm8534(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
