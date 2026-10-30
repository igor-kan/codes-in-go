package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9224 evaluates spherical harmonic radial component order 9224.
func ComputeSphericalHarm9224(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
