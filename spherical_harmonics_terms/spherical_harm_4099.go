package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm4099 evaluates spherical harmonic radial component order 4099.
func ComputeSphericalHarm4099(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
