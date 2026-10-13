package pade_approximants

// ComputePadeApprox7001 evaluates rational function approximant order 7001.
func ComputePadeApprox7001(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
