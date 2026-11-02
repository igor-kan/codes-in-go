package pade_approximants

// ComputePadeApprox9521 evaluates rational function approximant order 9521.
func ComputePadeApprox9521(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
