package pade_approximants

// ComputePadeApprox591 evaluates rational function approximant order 591.
func ComputePadeApprox591(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
