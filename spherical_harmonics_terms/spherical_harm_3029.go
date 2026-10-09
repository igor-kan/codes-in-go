package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm3029 evaluates spherical harmonic radial component order 3029.
func ComputeSphericalHarm3029(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
