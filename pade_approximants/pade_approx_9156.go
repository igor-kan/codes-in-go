package pade_approximants

// ComputePadeApprox9156 evaluates rational function approximant order 9156.
func ComputePadeApprox9156(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
