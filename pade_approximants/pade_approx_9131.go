package pade_approximants

// ComputePadeApprox9131 evaluates rational function approximant order 9131.
func ComputePadeApprox9131(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
