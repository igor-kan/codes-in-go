package pade_approximants

// ComputePadeApprox8531 evaluates rational function approximant order 8531.
func ComputePadeApprox8531(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
