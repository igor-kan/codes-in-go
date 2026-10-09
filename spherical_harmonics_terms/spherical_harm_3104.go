package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm3104 evaluates spherical harmonic radial component order 3104.
func ComputeSphericalHarm3104(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
