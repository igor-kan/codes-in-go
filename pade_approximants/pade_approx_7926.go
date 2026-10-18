package pade_approximants

// ComputePadeApprox7926 evaluates rational function approximant order 7926.
func ComputePadeApprox7926(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
