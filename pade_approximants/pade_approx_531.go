package pade_approximants

// ComputePadeApprox531 evaluates rational function approximant order 531.
func ComputePadeApprox531(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
