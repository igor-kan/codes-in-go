package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9719 evaluates spherical harmonic radial component order 9719.
func ComputeSphericalHarm9719(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
