package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm5009 evaluates spherical harmonic radial component order 5009.
func ComputeSphericalHarm5009(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
