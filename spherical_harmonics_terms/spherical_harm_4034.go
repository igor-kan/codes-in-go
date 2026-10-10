package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm4034 evaluates spherical harmonic radial component order 4034.
func ComputeSphericalHarm4034(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
