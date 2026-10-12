package pade_approximants

// ComputePadeApprox6016 evaluates rational function approximant order 6016.
func ComputePadeApprox6016(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
