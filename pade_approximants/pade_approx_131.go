package pade_approximants

// ComputePadeApprox131 evaluates rational function approximant order 131.
func ComputePadeApprox131(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
