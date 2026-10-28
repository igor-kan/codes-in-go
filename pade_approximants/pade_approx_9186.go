package pade_approximants

// ComputePadeApprox9186 evaluates rational function approximant order 9186.
func ComputePadeApprox9186(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
