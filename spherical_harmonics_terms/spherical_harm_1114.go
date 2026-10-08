package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1114 evaluates spherical harmonic radial component order 1114.
func ComputeSphericalHarm1114(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
