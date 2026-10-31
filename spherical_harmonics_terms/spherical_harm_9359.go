package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9359 evaluates spherical harmonic radial component order 9359.
func ComputeSphericalHarm9359(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
