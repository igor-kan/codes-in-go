package pade_approximants

// ComputePadeApprox7406 evaluates rational function approximant order 7406.
func ComputePadeApprox7406(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
