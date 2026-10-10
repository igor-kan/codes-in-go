package pade_approximants

// ComputePadeApprox4066 evaluates rational function approximant order 4066.
func ComputePadeApprox4066(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
