package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm4014 evaluates spherical harmonic radial component order 4014.
func ComputeSphericalHarm4014(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
