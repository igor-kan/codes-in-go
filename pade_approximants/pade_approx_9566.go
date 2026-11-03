package pade_approximants

// ComputePadeApprox9566 evaluates rational function approximant order 9566.
func ComputePadeApprox9566(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
