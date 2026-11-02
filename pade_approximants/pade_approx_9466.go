package pade_approximants

// ComputePadeApprox9466 evaluates rational function approximant order 9466.
func ComputePadeApprox9466(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
