package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm2014 evaluates spherical harmonic radial component order 2014.
func ComputeSphericalHarm2014(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
