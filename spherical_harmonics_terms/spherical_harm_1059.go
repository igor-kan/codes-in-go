package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1059 evaluates spherical harmonic radial component order 1059.
func ComputeSphericalHarm1059(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
