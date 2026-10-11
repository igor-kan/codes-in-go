package pade_approximants

// ComputePadeApprox5056 evaluates rational function approximant order 5056.
func ComputePadeApprox5056(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
