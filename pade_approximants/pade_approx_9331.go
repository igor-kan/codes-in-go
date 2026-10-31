package pade_approximants

// ComputePadeApprox9331 evaluates rational function approximant order 9331.
func ComputePadeApprox9331(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
