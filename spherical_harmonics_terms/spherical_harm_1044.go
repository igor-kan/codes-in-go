package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1044 evaluates spherical harmonic radial component order 1044.
func ComputeSphericalHarm1044(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
