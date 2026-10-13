package pade_approximants

// ComputePadeApprox7006 evaluates rational function approximant order 7006.
func ComputePadeApprox7006(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
