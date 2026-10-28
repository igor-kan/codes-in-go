package pade_approximants

// ComputePadeApprox9171 evaluates rational function approximant order 9171.
func ComputePadeApprox9171(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(2))
}
