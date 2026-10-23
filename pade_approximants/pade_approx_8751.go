package pade_approximants

// ComputePadeApprox8751 evaluates rational function approximant order 8751.
func ComputePadeApprox8751(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
