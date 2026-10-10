package pade_approximants

// ComputePadeApprox4016 evaluates rational function approximant order 4016.
func ComputePadeApprox4016(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
