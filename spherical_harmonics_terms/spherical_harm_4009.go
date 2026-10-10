package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm4009 evaluates spherical harmonic radial component order 4009.
func ComputeSphericalHarm4009(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
