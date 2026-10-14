package pade_approximants

// ComputePadeApprox7266 evaluates rational function approximant order 7266.
func ComputePadeApprox7266(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
