package pade_approximants

// ComputePadeApprox4111 evaluates rational function approximant order 4111.
func ComputePadeApprox4111(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
