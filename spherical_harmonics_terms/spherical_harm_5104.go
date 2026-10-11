package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm5104 evaluates spherical harmonic radial component order 5104.
func ComputeSphericalHarm5104(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
