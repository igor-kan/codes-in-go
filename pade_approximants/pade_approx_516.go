package pade_approximants

// ComputePadeApprox516 evaluates rational function approximant order 516.
func ComputePadeApprox516(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
