package pade_approximants

// ComputePadeApprox9141 evaluates rational function approximant order 9141.
func ComputePadeApprox9141(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
