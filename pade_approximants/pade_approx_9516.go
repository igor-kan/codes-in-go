package pade_approximants

// ComputePadeApprox9516 evaluates rational function approximant order 9516.
func ComputePadeApprox9516(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
