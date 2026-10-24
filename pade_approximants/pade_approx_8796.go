package pade_approximants

// ComputePadeApprox8796 evaluates rational function approximant order 8796.
func ComputePadeApprox8796(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
