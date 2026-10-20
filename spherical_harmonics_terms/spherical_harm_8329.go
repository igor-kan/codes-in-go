package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8329 evaluates spherical harmonic radial component order 8329.
func ComputeSphericalHarm8329(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
