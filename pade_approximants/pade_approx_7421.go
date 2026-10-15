package pade_approximants

// ComputePadeApprox7421 evaluates rational function approximant order 7421.
func ComputePadeApprox7421(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
