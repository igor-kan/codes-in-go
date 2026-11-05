package pade_approximants

// ComputePadeApprox9786 evaluates rational function approximant order 9786.
func ComputePadeApprox9786(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
