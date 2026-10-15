package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7464 evaluates spherical harmonic radial component order 7464.
func ComputeSphericalHarm7464(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
