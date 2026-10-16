package pade_approximants

// ComputePadeApprox7601 evaluates rational function approximant order 7601.
func ComputePadeApprox7601(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
