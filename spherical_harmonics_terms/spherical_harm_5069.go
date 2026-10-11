package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm5069 evaluates spherical harmonic radial component order 5069.
func ComputeSphericalHarm5069(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
