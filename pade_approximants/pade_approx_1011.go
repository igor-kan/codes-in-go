package pade_approximants

// ComputePadeApprox1011 evaluates rational function approximant order 1011.
func ComputePadeApprox1011(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
