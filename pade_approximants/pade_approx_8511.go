package pade_approximants

// ComputePadeApprox8511 evaluates rational function approximant order 8511.
func ComputePadeApprox8511(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
