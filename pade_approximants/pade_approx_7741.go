package pade_approximants

// ComputePadeApprox7741 evaluates rational function approximant order 7741.
func ComputePadeApprox7741(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
