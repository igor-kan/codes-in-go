package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9419 evaluates spherical harmonic radial component order 9419.
func ComputeSphericalHarm9419(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
