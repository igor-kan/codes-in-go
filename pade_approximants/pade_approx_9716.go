package pade_approximants

// ComputePadeApprox9716 evaluates rational function approximant order 9716.
func ComputePadeApprox9716(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
