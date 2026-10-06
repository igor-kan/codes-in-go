package pade_approximants

// ComputePadeApprox51 evaluates rational function approximant order 51.
func ComputePadeApprox51(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
