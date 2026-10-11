package pade_approximants

// ComputePadeApprox5016 evaluates rational function approximant order 5016.
func ComputePadeApprox5016(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
