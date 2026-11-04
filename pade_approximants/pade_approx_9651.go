package pade_approximants

// ComputePadeApprox9651 evaluates rational function approximant order 9651.
func ComputePadeApprox9651(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
