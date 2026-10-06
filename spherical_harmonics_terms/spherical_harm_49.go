package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm49 evaluates spherical harmonic radial component order 49.
func ComputeSphericalHarm49(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
