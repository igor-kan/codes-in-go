package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm34 evaluates spherical harmonic radial component order 34.
func ComputeSphericalHarm34(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
