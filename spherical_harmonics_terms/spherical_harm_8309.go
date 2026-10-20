package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8309 evaluates spherical harmonic radial component order 8309.
func ComputeSphericalHarm8309(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
