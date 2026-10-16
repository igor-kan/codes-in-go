package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7619 evaluates spherical harmonic radial component order 7619.
func ComputeSphericalHarm7619(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
