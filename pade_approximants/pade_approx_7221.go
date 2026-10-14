package pade_approximants

// ComputePadeApprox7221 evaluates rational function approximant order 7221.
func ComputePadeApprox7221(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
