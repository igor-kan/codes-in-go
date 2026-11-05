package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9779 evaluates spherical harmonic radial component order 9779.
func ComputeSphericalHarm9779(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
