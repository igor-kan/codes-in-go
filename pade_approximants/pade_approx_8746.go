package pade_approximants

// ComputePadeApprox8746 evaluates rational function approximant order 8746.
func ComputePadeApprox8746(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
