package pade_approximants

// ComputePadeApprox9626 evaluates rational function approximant order 9626.
func ComputePadeApprox9626(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
