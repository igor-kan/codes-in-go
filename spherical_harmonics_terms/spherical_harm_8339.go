package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm8339 evaluates spherical harmonic radial component order 8339.
func ComputeSphericalHarm8339(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
