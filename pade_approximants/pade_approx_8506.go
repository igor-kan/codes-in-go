package pade_approximants

// ComputePadeApprox8506 evaluates rational function approximant order 8506.
func ComputePadeApprox8506(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
