package pade_approximants

// ComputePadeApprox9096 evaluates rational function approximant order 9096.
func ComputePadeApprox9096(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
