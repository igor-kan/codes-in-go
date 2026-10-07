package pade_approximants

// ComputePadeApprox606 evaluates rational function approximant order 606.
func ComputePadeApprox606(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
