package pade_approximants

// ComputePadeApprox6026 evaluates rational function approximant order 6026.
func ComputePadeApprox6026(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
