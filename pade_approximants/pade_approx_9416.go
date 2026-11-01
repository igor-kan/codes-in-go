package pade_approximants

// ComputePadeApprox9416 evaluates rational function approximant order 9416.
func ComputePadeApprox9416(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
