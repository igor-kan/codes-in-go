package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7239 evaluates spherical harmonic radial component order 7239.
func ComputeSphericalHarm7239(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
