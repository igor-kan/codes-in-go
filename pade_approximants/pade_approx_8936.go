package pade_approximants

// ComputePadeApprox8936 evaluates rational function approximant order 8936.
func ComputePadeApprox8936(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
