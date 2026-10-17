package pade_approximants

// ComputePadeApprox7761 evaluates rational function approximant order 7761.
func ComputePadeApprox7761(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
