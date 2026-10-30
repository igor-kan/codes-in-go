package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9234 evaluates spherical harmonic radial component order 9234.
func ComputeSphericalHarm9234(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
