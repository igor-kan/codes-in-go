package pade_approximants

// ComputePadeApprox501 evaluates rational function approximant order 501.
func ComputePadeApprox501(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
