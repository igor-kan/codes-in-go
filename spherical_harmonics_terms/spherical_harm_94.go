package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm94 evaluates spherical harmonic radial component order 94.
func ComputeSphericalHarm94(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
