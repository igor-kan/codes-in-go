package pade_approximants

// ComputePadeApprox9221 evaluates rational function approximant order 9221.
func ComputePadeApprox9221(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
