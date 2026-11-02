package pade_approximants

// ComputePadeApprox9511 evaluates rational function approximant order 9511.
func ComputePadeApprox9511(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
