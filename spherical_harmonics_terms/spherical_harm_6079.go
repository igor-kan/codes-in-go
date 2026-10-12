package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm6079 evaluates spherical harmonic radial component order 6079.
func ComputeSphericalHarm6079(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
