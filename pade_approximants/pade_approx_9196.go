package pade_approximants

// ComputePadeApprox9196 evaluates rational function approximant order 9196.
func ComputePadeApprox9196(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
