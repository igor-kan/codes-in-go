package pade_approximants

// ComputePadeApprox8516 evaluates rational function approximant order 8516.
func ComputePadeApprox8516(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
