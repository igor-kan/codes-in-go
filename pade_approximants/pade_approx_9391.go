package pade_approximants

// ComputePadeApprox9391 evaluates rational function approximant order 9391.
func ComputePadeApprox9391(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
