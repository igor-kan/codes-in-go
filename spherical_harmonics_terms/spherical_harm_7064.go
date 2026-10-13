package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7064 evaluates spherical harmonic radial component order 7064.
func ComputeSphericalHarm7064(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
