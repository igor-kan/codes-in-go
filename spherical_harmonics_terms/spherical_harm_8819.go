package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8819 evaluates spherical harmonic radial component order 8819.
func ComputeSphericalHarm8819(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
