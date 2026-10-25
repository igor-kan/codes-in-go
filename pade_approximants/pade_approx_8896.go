package pade_approximants

// ComputePadeApprox8896 evaluates rational function approximant order 8896.
func ComputePadeApprox8896(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
