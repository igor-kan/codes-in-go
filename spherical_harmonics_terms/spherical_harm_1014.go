package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1014 evaluates spherical harmonic radial component order 1014.
func ComputeSphericalHarm1014(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
