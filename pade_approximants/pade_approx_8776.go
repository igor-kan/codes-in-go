package pade_approximants

// ComputePadeApprox8776 evaluates rational function approximant order 8776.
func ComputePadeApprox8776(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
