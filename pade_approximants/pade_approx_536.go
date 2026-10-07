package pade_approximants

// ComputePadeApprox536 evaluates rational function approximant order 536.
func ComputePadeApprox536(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
