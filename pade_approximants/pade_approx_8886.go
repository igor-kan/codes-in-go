package pade_approximants

// ComputePadeApprox8886 evaluates rational function approximant order 8886.
func ComputePadeApprox8886(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
