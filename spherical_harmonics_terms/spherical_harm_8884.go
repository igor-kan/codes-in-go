package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8884 evaluates spherical harmonic radial component order 8884.
func ComputeSphericalHarm8884(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
