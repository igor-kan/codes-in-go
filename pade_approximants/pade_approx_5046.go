package pade_approximants

// ComputePadeApprox5046 evaluates rational function approximant order 5046.
func ComputePadeApprox5046(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
