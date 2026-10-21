package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8504 evaluates spherical harmonic radial component order 8504.
func ComputeSphericalHarm8504(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
