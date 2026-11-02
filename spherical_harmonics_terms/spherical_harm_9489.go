package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9489 evaluates spherical harmonic radial component order 9489.
func ComputeSphericalHarm9489(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
