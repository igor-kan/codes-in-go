package pade_approximants

// ComputePadeApprox96 evaluates rational function approximant order 96.
func ComputePadeApprox96(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
