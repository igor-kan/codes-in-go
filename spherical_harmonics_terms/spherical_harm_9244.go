package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9244 evaluates spherical harmonic radial component order 9244.
func ComputeSphericalHarm9244(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
