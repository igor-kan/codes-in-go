package pade_approximants

// ComputePadeApprox581 evaluates rational function approximant order 581.
func ComputePadeApprox581(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
