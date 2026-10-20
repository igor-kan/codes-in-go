package pade_approximants

// ComputePadeApprox8336 evaluates rational function approximant order 8336.
func ComputePadeApprox8336(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
