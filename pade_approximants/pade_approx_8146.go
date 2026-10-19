package pade_approximants

// ComputePadeApprox8146 evaluates rational function approximant order 8146.
func ComputePadeApprox8146(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
