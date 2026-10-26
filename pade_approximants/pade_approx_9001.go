package pade_approximants

// ComputePadeApprox9001 evaluates rational function approximant order 9001.
func ComputePadeApprox9001(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
