package pade_approximants

// ComputePadeApprox4091 evaluates rational function approximant order 4091.
func ComputePadeApprox4091(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
