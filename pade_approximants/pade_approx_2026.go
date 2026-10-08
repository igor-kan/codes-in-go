package pade_approximants

// ComputePadeApprox2026 evaluates rational function approximant order 2026.
func ComputePadeApprox2026(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
