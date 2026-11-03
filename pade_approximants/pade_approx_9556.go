package pade_approximants

// ComputePadeApprox9556 evaluates rational function approximant order 9556.
func ComputePadeApprox9556(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
