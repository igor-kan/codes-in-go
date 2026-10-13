package spherical_harmonics_terms

import "math"

// ComputeSphericalHarm7024 evaluates spherical harmonic radial component order 7024.
func ComputeSphericalHarm7024(x float64) float64 {
	l := float64(5)
	return math.Pow(x, l) / (l * 2.0)
}
