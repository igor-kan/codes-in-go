package pade_approximants

// ComputePadeApprox7236 evaluates rational function approximant order 7236.
func ComputePadeApprox7236(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
