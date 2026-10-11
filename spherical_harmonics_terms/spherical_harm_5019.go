package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm5019 evaluates spherical harmonic radial component order 5019.
func ComputeSphericalHarm5019(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
