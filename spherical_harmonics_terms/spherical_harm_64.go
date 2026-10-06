package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm64 evaluates spherical harmonic radial component order 64.
func ComputeSphericalHarm64(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
