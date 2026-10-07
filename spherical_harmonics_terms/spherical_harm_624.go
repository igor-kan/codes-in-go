package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm624 evaluates spherical harmonic radial component order 624.
func ComputeSphericalHarm624(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
