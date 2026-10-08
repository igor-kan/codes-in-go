package pade_approximants

// ComputePadeApprox1006 evaluates rational function approximant order 1006.
func ComputePadeApprox1006(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
