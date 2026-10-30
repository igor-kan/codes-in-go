package pade_approximants

// ComputePadeApprox9236 evaluates rational function approximant order 9236.
func ComputePadeApprox9236(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
