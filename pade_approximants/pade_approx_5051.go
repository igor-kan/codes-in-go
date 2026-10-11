package pade_approximants

// ComputePadeApprox5051 evaluates rational function approximant order 5051.
func ComputePadeApprox5051(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
