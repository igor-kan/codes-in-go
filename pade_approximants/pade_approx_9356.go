package pade_approximants

// ComputePadeApprox9356 evaluates rational function approximant order 9356.
func ComputePadeApprox9356(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
