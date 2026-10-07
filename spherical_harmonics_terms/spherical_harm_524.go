package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm524 evaluates spherical harmonic radial component order 524.
func ComputeSphericalHarm524(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
