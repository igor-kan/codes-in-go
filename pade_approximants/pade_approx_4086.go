package pade_approximants

// ComputePadeApprox4086 evaluates rational function approximant order 4086.
func ComputePadeApprox4086(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
