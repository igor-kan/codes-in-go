package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm564 evaluates spherical harmonic radial component order 564.
func ComputeSphericalHarm564(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
