package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7424 evaluates spherical harmonic radial component order 7424.
func ComputeSphericalHarm7424(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
