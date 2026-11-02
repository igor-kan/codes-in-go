package pade_approximants

// ComputePadeApprox9491 evaluates rational function approximant order 9491.
func ComputePadeApprox9491(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
