package pade_approximants

// ComputePadeApprox9581 evaluates rational function approximant order 9581.
func ComputePadeApprox9581(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
