package pade_approximants

// ComputePadeApprox9381 evaluates rational function approximant order 9381.
func ComputePadeApprox9381(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
