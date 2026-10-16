package pade_approximants

// ComputePadeApprox7606 evaluates rational function approximant order 7606.
func ComputePadeApprox7606(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
