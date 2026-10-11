package pade_approximants

// ComputePadeApprox5106 evaluates rational function approximant order 5106.
func ComputePadeApprox5106(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
