package pade_approximants

// ComputePadeApprox7261 evaluates rational function approximant order 7261.
func ComputePadeApprox7261(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
