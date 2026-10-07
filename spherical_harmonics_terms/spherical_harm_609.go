package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm609 evaluates spherical harmonic radial component order 609.
func ComputeSphericalHarm609(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
