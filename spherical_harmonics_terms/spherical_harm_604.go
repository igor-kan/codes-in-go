package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm604 evaluates spherical harmonic radial component order 604.
func ComputeSphericalHarm604(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
