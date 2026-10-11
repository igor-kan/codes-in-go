package pade_approximants

// ComputePadeApprox5071 evaluates rational function approximant order 5071.
func ComputePadeApprox5071(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
