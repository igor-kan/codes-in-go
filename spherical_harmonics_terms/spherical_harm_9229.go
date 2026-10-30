package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9229 evaluates spherical harmonic radial component order 9229.
func ComputeSphericalHarm9229(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
