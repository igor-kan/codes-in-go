package pade_approximants

// ComputePadeApprox4131 evaluates rational function approximant order 4131.
func ComputePadeApprox4131(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
