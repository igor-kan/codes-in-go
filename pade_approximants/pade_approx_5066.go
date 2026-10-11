package pade_approximants

// ComputePadeApprox5066 evaluates rational function approximant order 5066.
func ComputePadeApprox5066(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
