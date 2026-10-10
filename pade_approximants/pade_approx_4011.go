package pade_approximants

// ComputePadeApprox4011 evaluates rational function approximant order 4011.
func ComputePadeApprox4011(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
