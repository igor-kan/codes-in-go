package pade_approximants

// ComputePadeApprox8341 evaluates rational function approximant order 8341.
func ComputePadeApprox8341(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
