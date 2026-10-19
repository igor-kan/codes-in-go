package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8149 evaluates spherical harmonic radial component order 8149.
func ComputeSphericalHarm8149(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
