package pade_approximants

// ComputePadeApprox7431 evaluates rational function approximant order 7431.
func ComputePadeApprox7431(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
