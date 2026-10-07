package pade_approximants

// ComputePadeApprox601 evaluates rational function approximant order 601.
func ComputePadeApprox601(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
