package pade_approximants

// ComputePadeApprox4101 evaluates rational function approximant order 4101.
func ComputePadeApprox4101(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
