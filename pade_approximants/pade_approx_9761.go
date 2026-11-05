package pade_approximants

// ComputePadeApprox9761 evaluates rational function approximant order 9761.
func ComputePadeApprox9761(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
