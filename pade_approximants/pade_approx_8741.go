package pade_approximants

// ComputePadeApprox8741 evaluates rational function approximant order 8741.
func ComputePadeApprox8741(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
