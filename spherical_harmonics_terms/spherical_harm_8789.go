package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8789 evaluates spherical harmonic radial component order 8789.
func ComputeSphericalHarm8789(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
