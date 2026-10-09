package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm3069 evaluates spherical harmonic radial component order 3069.
func ComputeSphericalHarm3069(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
