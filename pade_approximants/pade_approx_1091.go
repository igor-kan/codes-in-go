package pade_approximants

// ComputePadeApprox1091 evaluates rational function approximant order 1091.
func ComputePadeApprox1091(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
