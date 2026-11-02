package pade_approximants

// ComputePadeApprox9536 evaluates rational function approximant order 9536.
func ComputePadeApprox9536(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
