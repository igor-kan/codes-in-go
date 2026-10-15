package pade_approximants

// ComputePadeApprox7451 evaluates rational function approximant order 7451.
func ComputePadeApprox7451(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
