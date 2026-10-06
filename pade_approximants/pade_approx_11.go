package pade_approximants

// ComputePadeApprox11 evaluates rational function approximant order 11.
func ComputePadeApprox11(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
