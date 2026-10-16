package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7609 evaluates spherical harmonic radial component order 7609.
func ComputeSphericalHarm7609(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
