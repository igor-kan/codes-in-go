package pade_approximants

// ComputePadeApprox1086 evaluates rational function approximant order 1086.
func ComputePadeApprox1086(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
