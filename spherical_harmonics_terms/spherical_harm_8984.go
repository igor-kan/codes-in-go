package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8984 evaluates spherical harmonic radial component order 8984.
func ComputeSphericalHarm8984(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
