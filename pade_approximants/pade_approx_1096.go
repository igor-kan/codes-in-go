package pade_approximants

// ComputePadeApprox1096 evaluates rational function approximant order 1096.
func ComputePadeApprox1096(x float64) float64 {
	return (1.0 + x*float64(2)) / (1.0 + x*x*float64(1))
}
