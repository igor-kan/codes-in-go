package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9404 evaluates spherical harmonic radial component order 9404.
func ComputeSphericalHarm9404(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
