package pade_approximants

// ComputePadeApprox8921 evaluates rational function approximant order 8921.
func ComputePadeApprox8921(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(2))
}
