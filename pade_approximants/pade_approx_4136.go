package pade_approximants

// ComputePadeApprox4136 evaluates rational function approximant order 4136.
func ComputePadeApprox4136(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
