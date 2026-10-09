package pade_approximants

// ComputePadeApprox3091 evaluates rational function approximant order 3091.
func ComputePadeApprox3091(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
