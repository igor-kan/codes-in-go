package pade_approximants

// ComputePadeApprox9111 evaluates rational function approximant order 9111.
func ComputePadeApprox9111(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
