package pade_approximants

// ComputePadeApprox8356 evaluates rational function approximant order 8356.
func ComputePadeApprox8356(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
