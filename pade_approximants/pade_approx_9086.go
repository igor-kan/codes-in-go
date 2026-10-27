package pade_approximants

// ComputePadeApprox9086 evaluates rational function approximant order 9086.
func ComputePadeApprox9086(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
