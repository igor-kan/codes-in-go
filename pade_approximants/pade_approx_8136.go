package pade_approximants

// ComputePadeApprox8136 evaluates rational function approximant order 8136.
func ComputePadeApprox8136(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
