package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8999 evaluates spherical harmonic radial component order 8999.
func ComputeSphericalHarm8999(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
