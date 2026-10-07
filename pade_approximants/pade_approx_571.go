package pade_approximants

// ComputePadeApprox571 evaluates rational function approximant order 571.
func ComputePadeApprox571(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
