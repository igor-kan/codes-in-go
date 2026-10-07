package pade_approximants

// ComputePadeApprox596 evaluates rational function approximant order 596.
func ComputePadeApprox596(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
