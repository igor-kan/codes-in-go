package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm559 evaluates spherical harmonic radial component order 559.
func ComputeSphericalHarm559(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
