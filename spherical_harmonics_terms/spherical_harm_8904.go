package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8904 evaluates spherical harmonic radial component order 8904.
func ComputeSphericalHarm8904(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
