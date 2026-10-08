package pade_approximants

// ComputePadeApprox1026 evaluates rational function approximant order 1026.
func ComputePadeApprox1026(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
