package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7949 evaluates spherical harmonic radial component order 7949.
func ComputeSphericalHarm7949(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
