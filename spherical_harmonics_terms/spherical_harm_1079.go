package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1079 evaluates spherical harmonic radial component order 1079.
func ComputeSphericalHarm1079(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
