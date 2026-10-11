package pade_approximants

// ComputePadeApprox5061 evaluates rational function approximant order 5061.
func ComputePadeApprox5061(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
