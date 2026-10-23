package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8749 evaluates spherical harmonic radial component order 8749.
func ComputeSphericalHarm8749(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
