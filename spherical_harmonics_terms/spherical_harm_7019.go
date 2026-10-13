package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7019 evaluates spherical harmonic radial component order 7019.
func ComputeSphericalHarm7019(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
