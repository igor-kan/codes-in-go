package pade_approximants

// ComputePadeApprox4006 evaluates rational function approximant order 4006.
func ComputePadeApprox4006(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
