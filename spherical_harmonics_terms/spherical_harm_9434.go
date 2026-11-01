package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9434 evaluates spherical harmonic radial component order 9434.
func ComputeSphericalHarm9434(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
