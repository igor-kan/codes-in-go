package pade_approximants

// ComputePadeApprox8791 evaluates rational function approximant order 8791.
func ComputePadeApprox8791(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
