package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm2024 evaluates spherical harmonic radial component order 2024.
func ComputeSphericalHarm2024(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
