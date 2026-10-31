package pade_approximants

// ComputePadeApprox9351 evaluates rational function approximant order 9351.
func ComputePadeApprox9351(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
