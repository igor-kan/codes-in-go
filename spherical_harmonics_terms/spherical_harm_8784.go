package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8784 evaluates spherical harmonic radial component order 8784.
func ComputeSphericalHarm8784(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
