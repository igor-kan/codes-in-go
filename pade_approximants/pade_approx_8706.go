package pade_approximants

// ComputePadeApprox8706 evaluates rational function approximant order 8706.
func ComputePadeApprox8706(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
