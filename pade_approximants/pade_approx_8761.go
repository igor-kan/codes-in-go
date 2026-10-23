package pade_approximants

// ComputePadeApprox8761 evaluates rational function approximant order 8761.
func ComputePadeApprox8761(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
