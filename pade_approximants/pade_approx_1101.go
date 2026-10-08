package pade_approximants

// ComputePadeApprox1101 evaluates rational function approximant order 1101.
func ComputePadeApprox1101(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
