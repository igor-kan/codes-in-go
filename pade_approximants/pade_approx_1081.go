package pade_approximants

// ComputePadeApprox1081 evaluates rational function approximant order 1081.
func ComputePadeApprox1081(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
