package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm9199 evaluates spherical harmonic radial component order 9199.
func ComputeSphericalHarm9199(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
