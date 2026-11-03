package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9589 evaluates spherical harmonic radial component order 9589.
func ComputeSphericalHarm9589(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
