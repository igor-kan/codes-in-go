package pade_approximants

// ComputePadeApprox8101 evaluates rational function approximant order 8101.
func ComputePadeApprox8101(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
