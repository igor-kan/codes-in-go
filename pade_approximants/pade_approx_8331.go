package pade_approximants

// ComputePadeApprox8331 evaluates rational function approximant order 8331.
func ComputePadeApprox8331(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
