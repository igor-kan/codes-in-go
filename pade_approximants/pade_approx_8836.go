package pade_approximants

// ComputePadeApprox8836 evaluates rational function approximant order 8836.
func ComputePadeApprox8836(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
