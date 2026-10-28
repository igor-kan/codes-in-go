package pade_approximants

// ComputePadeApprox9201 evaluates rational function approximant order 9201.
func ComputePadeApprox9201(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
