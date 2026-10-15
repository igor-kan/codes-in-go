package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7444 evaluates spherical harmonic radial component order 7444.
func ComputeSphericalHarm7444(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
