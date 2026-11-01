package pade_approximants

// ComputePadeApprox9446 evaluates rational function approximant order 9446.
func ComputePadeApprox9446(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
