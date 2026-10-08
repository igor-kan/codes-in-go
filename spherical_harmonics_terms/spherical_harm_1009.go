package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1009 evaluates spherical harmonic radial component order 1009.
func ComputeSphericalHarm1009(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
