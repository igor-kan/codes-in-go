package pade_approximants

// ComputePadeApprox7461 evaluates rational function approximant order 7461.
func ComputePadeApprox7461(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
