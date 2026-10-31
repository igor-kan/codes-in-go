package pade_approximants

// ComputePadeApprox9316 evaluates rational function approximant order 9316.
func ComputePadeApprox9316(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
