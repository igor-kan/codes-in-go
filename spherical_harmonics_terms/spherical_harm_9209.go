package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9209 evaluates spherical harmonic radial component order 9209.
func ComputeSphericalHarm9209(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
