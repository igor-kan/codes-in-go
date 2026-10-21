package pade_approximants

// ComputePadeApprox8556 evaluates rational function approximant order 8556.
func ComputePadeApprox8556(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
