package pade_approximants

// ComputePadeApprox8946 evaluates rational function approximant order 8946.
func ComputePadeApprox8946(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
