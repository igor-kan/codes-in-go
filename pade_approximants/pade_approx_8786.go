package pade_approximants

// ComputePadeApprox8786 evaluates rational function approximant order 8786.
func ComputePadeApprox8786(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
