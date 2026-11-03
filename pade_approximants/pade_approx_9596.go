package pade_approximants

// ComputePadeApprox9596 evaluates rational function approximant order 9596.
func ComputePadeApprox9596(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
