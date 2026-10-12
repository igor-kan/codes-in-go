package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm6009 evaluates spherical harmonic radial component order 6009.
func ComputeSphericalHarm6009(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
