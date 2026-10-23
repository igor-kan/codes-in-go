package pade_approximants

// ComputePadeApprox8756 evaluates rational function approximant order 8756.
func ComputePadeApprox8756(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
