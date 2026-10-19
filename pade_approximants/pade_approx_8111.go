package pade_approximants

// ComputePadeApprox8111 evaluates rational function approximant order 8111.
func ComputePadeApprox8111(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
