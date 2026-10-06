package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm79 evaluates spherical harmonic radial component order 79.
func ComputeSphericalHarm79(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
