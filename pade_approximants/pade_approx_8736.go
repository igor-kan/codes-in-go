package pade_approximants

// ComputePadeApprox8736 evaluates rational function approximant order 8736.
func ComputePadeApprox8736(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
