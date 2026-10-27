package pade_approximants

// ComputePadeApprox9136 evaluates rational function approximant order 9136.
func ComputePadeApprox9136(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
