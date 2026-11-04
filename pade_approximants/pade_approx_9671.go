package pade_approximants

// ComputePadeApprox9671 evaluates rational function approximant order 9671.
func ComputePadeApprox9671(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
