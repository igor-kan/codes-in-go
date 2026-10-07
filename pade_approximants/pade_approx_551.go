package pade_approximants

// ComputePadeApprox551 evaluates rational function approximant order 551.
func ComputePadeApprox551(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
