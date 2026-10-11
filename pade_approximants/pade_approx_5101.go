package pade_approximants

// ComputePadeApprox5101 evaluates rational function approximant order 5101.
func ComputePadeApprox5101(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
