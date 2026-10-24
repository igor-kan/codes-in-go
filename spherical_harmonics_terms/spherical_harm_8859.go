package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8859 evaluates spherical harmonic radial component order 8859.
func ComputeSphericalHarm8859(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
