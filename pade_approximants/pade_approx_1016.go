package pade_approximants

// ComputePadeApprox1016 evaluates rational function approximant order 1016.
func ComputePadeApprox1016(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
