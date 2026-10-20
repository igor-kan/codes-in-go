package pade_approximants

// ComputePadeApprox8351 evaluates rational function approximant order 8351.
func ComputePadeApprox8351(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
