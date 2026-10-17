package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7739 evaluates spherical harmonic radial component order 7739.
func ComputeSphericalHarm7739(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
