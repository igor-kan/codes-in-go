package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9214 evaluates spherical harmonic radial component order 9214.
func ComputeSphericalHarm9214(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
