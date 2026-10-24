package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8779 evaluates spherical harmonic radial component order 8779.
func ComputeSphericalHarm8779(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
