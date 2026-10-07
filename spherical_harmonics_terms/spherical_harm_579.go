package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm579 evaluates spherical harmonic radial component order 579.
func ComputeSphericalHarm579(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
