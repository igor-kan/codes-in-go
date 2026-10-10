package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm4064 evaluates spherical harmonic radial component order 4064.
func ComputeSphericalHarm4064(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
