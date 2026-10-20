package pade_approximants

// ComputePadeApprox8316 evaluates rational function approximant order 8316.
func ComputePadeApprox8316(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
