package pade_approximants

// ComputePadeApprox9776 evaluates rational function approximant order 9776.
func ComputePadeApprox9776(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
