package pade_approximants

// ComputePadeApprox4046 evaluates rational function approximant order 4046.
func ComputePadeApprox4046(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
