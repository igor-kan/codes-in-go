package pade_approximants

// ComputePadeApprox1001 evaluates rational function approximant order 1001.
func ComputePadeApprox1001(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
