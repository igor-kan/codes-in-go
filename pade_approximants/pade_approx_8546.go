package pade_approximants

// ComputePadeApprox8546 evaluates rational function approximant order 8546.
func ComputePadeApprox8546(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
