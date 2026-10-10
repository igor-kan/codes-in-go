package pade_approximants

// ComputePadeApprox4141 evaluates rational function approximant order 4141.
func ComputePadeApprox4141(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
