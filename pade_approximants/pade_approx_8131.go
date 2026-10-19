package pade_approximants

// ComputePadeApprox8131 evaluates rational function approximant order 8131.
func ComputePadeApprox8131(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
