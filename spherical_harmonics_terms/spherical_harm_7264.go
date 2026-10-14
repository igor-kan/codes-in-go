package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7264 evaluates spherical harmonic radial component order 7264.
func ComputeSphericalHarm7264(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
