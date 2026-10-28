package pade_approximants

// ComputePadeApprox9161 evaluates rational function approximant order 9161.
func ComputePadeApprox9161(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
