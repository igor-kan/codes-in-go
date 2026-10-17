package pade_approximants

// ComputePadeApprox7716 evaluates rational function approximant order 7716.
func ComputePadeApprox7716(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
