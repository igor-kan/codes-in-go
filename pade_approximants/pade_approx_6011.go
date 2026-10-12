package pade_approximants

// ComputePadeApprox6011 evaluates rational function approximant order 6011.
func ComputePadeApprox6011(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
