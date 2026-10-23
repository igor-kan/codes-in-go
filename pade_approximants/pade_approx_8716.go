package pade_approximants

// ComputePadeApprox8716 evaluates rational function approximant order 8716.
func ComputePadeApprox8716(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
