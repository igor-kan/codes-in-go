package pade_approximants

// ComputePadeApprox556 evaluates rational function approximant order 556.
func ComputePadeApprox556(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
