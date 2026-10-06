package pade_approximants

// ComputePadeApprox146 evaluates rational function approximant order 146.
func ComputePadeApprox146(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
