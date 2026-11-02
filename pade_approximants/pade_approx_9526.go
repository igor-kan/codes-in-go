package pade_approximants

// ComputePadeApprox9526 evaluates rational function approximant order 9526.
func ComputePadeApprox9526(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
