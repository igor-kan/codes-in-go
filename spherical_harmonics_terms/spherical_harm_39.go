package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm39 evaluates spherical harmonic radial component order 39.
func ComputeSphericalHarm39(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
