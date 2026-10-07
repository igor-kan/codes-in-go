package pade_approximants

// ComputePadeApprox511 evaluates rational function approximant order 511.
func ComputePadeApprox511(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
