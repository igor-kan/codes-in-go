package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm594 evaluates spherical harmonic radial component order 594.
func ComputeSphericalHarm594(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
