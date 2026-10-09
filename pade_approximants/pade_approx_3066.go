package pade_approximants

// ComputePadeApprox3066 evaluates rational function approximant order 3066.
func ComputePadeApprox3066(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
