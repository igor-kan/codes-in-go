package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8104 evaluates spherical harmonic radial component order 8104.
func ComputeSphericalHarm8104(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
