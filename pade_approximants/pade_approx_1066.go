package pade_approximants

// ComputePadeApprox1066 evaluates rational function approximant order 1066.
func ComputePadeApprox1066(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
