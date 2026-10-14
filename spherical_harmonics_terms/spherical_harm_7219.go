package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7219 evaluates spherical harmonic radial component order 7219.
func ComputeSphericalHarm7219(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
