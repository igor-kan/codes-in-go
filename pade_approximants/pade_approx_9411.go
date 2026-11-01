package pade_approximants

// ComputePadeApprox9411 evaluates rational function approximant order 9411.
func ComputePadeApprox9411(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
