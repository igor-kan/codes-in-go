package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm544 evaluates spherical harmonic radial component order 544.
func ComputeSphericalHarm544(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
