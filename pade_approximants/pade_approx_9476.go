package pade_approximants

// ComputePadeApprox9476 evaluates rational function approximant order 9476.
func ComputePadeApprox9476(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
