package pade_approximants

// ComputePadeApprox7711 evaluates rational function approximant order 7711.
func ComputePadeApprox7711(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
