package pade_approximants

// ComputePadeApprox5096 evaluates rational function approximant order 5096.
func ComputePadeApprox5096(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
