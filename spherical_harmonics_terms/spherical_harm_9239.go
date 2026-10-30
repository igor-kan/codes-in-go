package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9239 evaluates spherical harmonic radial component order 9239.
func ComputeSphericalHarm9239(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
