package pade_approximants

// ComputePadeApprox8856 evaluates rational function approximant order 8856.
func ComputePadeApprox8856(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
