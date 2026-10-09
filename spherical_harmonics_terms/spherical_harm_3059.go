package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm3059 evaluates spherical harmonic radial component order 3059.
func ComputeSphericalHarm3059(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
