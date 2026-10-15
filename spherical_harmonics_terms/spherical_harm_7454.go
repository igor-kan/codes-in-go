package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7454 evaluates spherical harmonic radial component order 7454.
func ComputeSphericalHarm7454(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
