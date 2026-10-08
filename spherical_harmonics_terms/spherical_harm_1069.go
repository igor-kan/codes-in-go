package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1069 evaluates spherical harmonic radial component order 1069.
func ComputeSphericalHarm1069(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
