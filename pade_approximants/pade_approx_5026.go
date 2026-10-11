package pade_approximants

// ComputePadeApprox5026 evaluates rational function approximant order 5026.
func ComputePadeApprox5026(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
