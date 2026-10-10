package pade_approximants

// ComputePadeApprox4001 evaluates rational function approximant order 4001.
func ComputePadeApprox4001(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
