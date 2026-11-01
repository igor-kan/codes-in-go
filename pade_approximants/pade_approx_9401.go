package pade_approximants

// ComputePadeApprox9401 evaluates rational function approximant order 9401.
func ComputePadeApprox9401(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
