package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7704 evaluates spherical harmonic radial component order 7704.
func ComputeSphericalHarm7704(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
