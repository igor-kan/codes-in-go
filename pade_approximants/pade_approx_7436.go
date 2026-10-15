package pade_approximants

// ComputePadeApprox7436 evaluates rational function approximant order 7436.
func ComputePadeApprox7436(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
