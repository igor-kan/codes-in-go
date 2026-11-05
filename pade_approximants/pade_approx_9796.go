package pade_approximants

// ComputePadeApprox9796 evaluates rational function approximant order 9796.
func ComputePadeApprox9796(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
