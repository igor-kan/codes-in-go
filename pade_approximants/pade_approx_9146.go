package pade_approximants

// ComputePadeApprox9146 evaluates rational function approximant order 9146.
func ComputePadeApprox9146(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
