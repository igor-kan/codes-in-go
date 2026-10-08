package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm1019 evaluates spherical harmonic radial component order 1019.
func ComputeSphericalHarm1019(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
