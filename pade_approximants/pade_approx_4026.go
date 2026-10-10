package pade_approximants

// ComputePadeApprox4026 evaluates rational function approximant order 4026.
func ComputePadeApprox4026(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
