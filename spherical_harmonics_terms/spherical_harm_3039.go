package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm3039 evaluates spherical harmonic radial component order 3039.
func ComputeSphericalHarm3039(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
