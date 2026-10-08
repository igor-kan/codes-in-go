package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1039 evaluates spherical harmonic radial component order 1039.
func ComputeSphericalHarm1039(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
