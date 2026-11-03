package pade_approximants

// ComputePadeApprox9571 evaluates rational function approximant order 9571.
func ComputePadeApprox9571(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
