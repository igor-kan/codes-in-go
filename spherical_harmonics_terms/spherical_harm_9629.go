package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9629 evaluates spherical harmonic radial component order 9629.
func ComputeSphericalHarm9629(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
