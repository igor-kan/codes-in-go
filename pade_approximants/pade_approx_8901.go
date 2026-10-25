package pade_approximants

// ComputePadeApprox8901 evaluates rational function approximant order 8901.
func ComputePadeApprox8901(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
