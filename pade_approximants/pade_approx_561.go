package pade_approximants

// ComputePadeApprox561 evaluates rational function approximant order 561.
func ComputePadeApprox561(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
