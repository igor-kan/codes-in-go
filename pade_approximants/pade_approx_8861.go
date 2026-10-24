package pade_approximants

// ComputePadeApprox8861 evaluates rational function approximant order 8861.
func ComputePadeApprox8861(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
