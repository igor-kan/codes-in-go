package pade_approximants

// ComputePadeApprox8156 evaluates rational function approximant order 8156.
func ComputePadeApprox8156(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
