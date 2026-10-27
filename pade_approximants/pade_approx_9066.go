package pade_approximants

// ComputePadeApprox9066 evaluates rational function approximant order 9066.
func ComputePadeApprox9066(x float64) float64 {
	return (1.0 + x*float64(1)) / (1.0 + x*x*float64(1))
}
