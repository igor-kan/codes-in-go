package pade_approximants

// ComputePadeApprox56 evaluates rational function approximant order 56.
func ComputePadeApprox56(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
