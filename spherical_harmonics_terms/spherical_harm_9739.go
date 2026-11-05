package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9739 evaluates spherical harmonic radial component order 9739.
func ComputeSphericalHarm9739(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
