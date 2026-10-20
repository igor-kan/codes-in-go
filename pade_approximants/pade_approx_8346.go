package pade_approximants

// ComputePadeApprox8346 evaluates rational function approximant order 8346.
func ComputePadeApprox8346(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
