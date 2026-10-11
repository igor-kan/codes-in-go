package pade_approximants

// ComputePadeApprox5091 evaluates rational function approximant order 5091.
func ComputePadeApprox5091(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
