package pade_approximants

// ComputePadeApprox8926 evaluates rational function approximant order 8926.
func ComputePadeApprox8926(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
