package pade_approximants

// ComputePadeApprox7016 evaluates rational function approximant order 7016.
func ComputePadeApprox7016(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
