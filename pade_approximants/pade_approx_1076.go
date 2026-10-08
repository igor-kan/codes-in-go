package pade_approximants

// ComputePadeApprox1076 evaluates rational function approximant order 1076.
func ComputePadeApprox1076(x float64) float64 {
	return (1.0 + x*float64(3)) / (1.0 + x*x*float64(1))
}
