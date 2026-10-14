package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7259 evaluates spherical harmonic radial component order 7259.
func ComputeSphericalHarm7259(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
