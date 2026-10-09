package pade_approximants

// ComputePadeApprox3111 evaluates rational function approximant order 3111.
func ComputePadeApprox3111(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
