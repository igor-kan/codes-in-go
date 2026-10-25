package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8934 evaluates spherical harmonic radial component order 8934.
func ComputeSphericalHarm8934(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
