package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm114 evaluates spherical harmonic radial component order 114.
func ComputeSphericalHarm114(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
