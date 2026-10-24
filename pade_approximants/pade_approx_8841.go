package pade_approximants

// ComputePadeApprox8841 evaluates rational function approximant order 8841.
func ComputePadeApprox8841(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
