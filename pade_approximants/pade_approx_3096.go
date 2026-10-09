package pade_approximants

// ComputePadeApprox3096 evaluates rational function approximant order 3096.
func ComputePadeApprox3096(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
