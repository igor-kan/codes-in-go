package pade_approximants

// ComputePadeApprox8871 evaluates rational function approximant order 8871.
func ComputePadeApprox8871(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
