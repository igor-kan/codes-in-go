package pade_approximants

// ComputePadeApprox1021 evaluates rational function approximant order 1021.
func ComputePadeApprox1021(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(2))
}
