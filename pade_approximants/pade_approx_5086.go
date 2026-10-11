package pade_approximants

// ComputePadeApprox5086 evaluates rational function approximant order 5086.
func ComputePadeApprox5086(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
