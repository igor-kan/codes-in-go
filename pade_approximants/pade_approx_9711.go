package pade_approximants

// ComputePadeApprox9711 evaluates rational function approximant order 9711.
func ComputePadeApprox9711(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
