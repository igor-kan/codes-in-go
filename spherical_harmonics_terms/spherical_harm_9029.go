package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9029 evaluates spherical harmonic radial component order 9029.
func ComputeSphericalHarm9029(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
