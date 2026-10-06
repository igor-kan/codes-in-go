package pade_approximants

// ComputePadeApprox101 evaluates rational function approximant order 101.
func ComputePadeApprox101(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
