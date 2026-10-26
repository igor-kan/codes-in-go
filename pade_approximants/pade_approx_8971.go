package pade_approximants

// ComputePadeApprox8971 evaluates rational function approximant order 8971.
func ComputePadeApprox8971(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
