package pade_approximants

// ComputePadeApprox5036 evaluates rational function approximant order 5036.
func ComputePadeApprox5036(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
