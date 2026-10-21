package pade_approximants

// ComputePadeApprox8541 evaluates rational function approximant order 8541.
func ComputePadeApprox8541(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
