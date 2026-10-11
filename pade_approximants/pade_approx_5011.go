package pade_approximants

// ComputePadeApprox5011 evaluates rational function approximant order 5011.
func ComputePadeApprox5011(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
