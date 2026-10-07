package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm584 evaluates spherical harmonic radial component order 584.
func ComputeSphericalHarm584(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
