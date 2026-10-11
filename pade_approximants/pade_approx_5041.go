package pade_approximants

// ComputePadeApprox5041 evaluates rational function approximant order 5041.
func ComputePadeApprox5041(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
