package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8524 evaluates spherical harmonic radial component order 8524.
func ComputeSphericalHarm8524(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
