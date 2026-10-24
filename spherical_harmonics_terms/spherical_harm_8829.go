package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8829 evaluates spherical harmonic radial component order 8829.
func ComputeSphericalHarm8829(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
