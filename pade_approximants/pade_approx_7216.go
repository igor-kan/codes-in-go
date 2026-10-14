package pade_approximants

// ComputePadeApprox7216 evaluates rational function approximant order 7216.
func ComputePadeApprox7216(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
