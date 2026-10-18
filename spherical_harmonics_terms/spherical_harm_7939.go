package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7939 evaluates spherical harmonic radial component order 7939.
func ComputeSphericalHarm7939(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
