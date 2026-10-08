package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1084 evaluates spherical harmonic radial component order 1084.
func ComputeSphericalHarm1084(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
