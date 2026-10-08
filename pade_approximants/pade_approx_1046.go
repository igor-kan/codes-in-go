package pade_approximants

// ComputePadeApprox1046 evaluates rational function approximant order 1046.
func ComputePadeApprox1046(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
