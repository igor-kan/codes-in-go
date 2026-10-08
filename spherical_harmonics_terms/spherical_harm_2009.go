package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm2009 evaluates spherical harmonic radial component order 2009.
func ComputeSphericalHarm2009(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
