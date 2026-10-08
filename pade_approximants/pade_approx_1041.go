package pade_approximants

// ComputePadeApprox1041 evaluates rational function approximant order 1041.
func ComputePadeApprox1041(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
