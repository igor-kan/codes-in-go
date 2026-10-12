package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm6014 evaluates spherical harmonic radial component order 6014.
func ComputeSphericalHarm6014(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
