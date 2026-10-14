package pade_approximants

// ComputePadeApprox7206 evaluates rational function approximant order 7206.
func ComputePadeApprox7206(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
