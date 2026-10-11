package pade_approximants

// ComputePadeApprox5001 evaluates rational function approximant order 5001.
func ComputePadeApprox5001(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
