package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8874 evaluates spherical harmonic radial component order 8874.
func ComputeSphericalHarm8874(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
