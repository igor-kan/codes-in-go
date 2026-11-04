package pade_approximants

// ComputePadeApprox9666 evaluates rational function approximant order 9666.
func ComputePadeApprox9666(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
