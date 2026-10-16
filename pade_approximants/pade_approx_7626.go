package pade_approximants

// ComputePadeApprox7626 evaluates rational function approximant order 7626.
func ComputePadeApprox7626(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
